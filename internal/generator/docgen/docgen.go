package docgen

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/octo-cli/octo-cli/internal/generator/supported"
	"github.com/octo-cli/octo-cli/internal/generator/swaggerparser"
	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/spf13/afero"
)

type paramHelp struct {
	name        string
	required    bool
	description string
	unsupported bool
}

type unsupportedOptionalParam struct {
	operation *openapi3.Operation
	routePath string
	method    string
	paramName string
}

func writeMarkdownHeader(w io.Writer, level int, val string) (int, error) {
	val = "\n" + strings.Repeat("#", level) + " " + val + "\n\n"
	return w.Write([]byte(val))
}

func writerMust(_ int, err error) {
	if err != nil {
		panic(err)
	}
}

func sortUnsupportedOptionalParams(vals []unsupportedOptionalParam) {
	sort.Slice(vals, func(i, j int) bool {
		if vals[i].operation.OperationID < vals[j].operation.OperationID {
			return true
		}
		if vals[i].operation.OperationID > vals[j].operation.OperationID {
			return false
		}
		return vals[i].paramName < vals[j].paramName
	})
}

func getUnsupportedOptionalParams(swagger *openapi3.Swagger) []unsupportedOptionalParam {
	var result []unsupportedOptionalParam
	swaggerparser.ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		if supported.OperationIsUnsupported(op, path, method) {
			return
		}
		if op.RequestBody == nil {
			return
		}
		content := op.RequestBody.Value.Content.Get("application/json")
		if content == nil {
			return
		}
		for name, ref := range content.Schema.Value.Properties {
			if supported.IsSupportedParam(ref) {
				continue
			}
			result = append(result, unsupportedOptionalParam{
				operation: op,
				routePath: path,
				method:    method,
				paramName: name,
			})
		}
	})
	sortUnsupportedOptionalParams(result)
	return result
}

func WriteDocs(routesPath, docsPath string, fs afero.Fs) error {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(routesPath)
	if err != nil {
		return err
	}
	err = swaggerparser.RemoveOwnerParams(swagger)
	if err != nil {
		return err
	}
	opDoc, err := operationsHelp(swagger)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, filepath.Join(docsPath, "operations.md"), opDoc, 0640)
	if err != nil {
		return err
	}
	unsupDoc, err := unsupportedHelp(swagger)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, filepath.Join(docsPath, "unsupported.md"), unsupDoc, 0640)
	if err != nil {
		return err
	}
	return nil
}

func operationsHelp(swagger *openapi3.Swagger) ([]byte, error) {
	var buf bytes.Buffer

	uops := getUnsupportedOptionalParams(swagger)

	for _, svcName := range swaggerparser.AllServiceNames(swagger) {
		svcOps := supported.SvcOperations(swagger, svcName)
		if len(svcOps) == 0 {
			continue
		}
		sort.Slice(svcOps, func(i, j int) bool {
			return svcOps[i].Op.OperationID < svcOps[j].Op.OperationID
		})
		writerMust(
			writeMarkdownHeader(&buf, 1, svcName),
		)
		for _, opInfo := range svcOps {
			writerMust(writeMarkdownHeader(&buf, 2, fmt.Sprintf("%s %s", svcName, swaggerparser.GetOperationName(opInfo.Op))))
			writerMust(buf.WriteString(opInfo.Op.ExternalDocs.URL + "\n\n"))
			writerMust(buf.WriteString(opInfo.Op.Description + "\n"))
			helps := paramHelps(opInfo.Path, opInfo.Method, opInfo.Op, uops)
			if len(helps) == 0 {
				continue
			}
			writerMust(writeMarkdownHeader(&buf, 3, "parameters"))

			writerMust(buf.WriteString(`
| name | description |
|------|-------------|
`))

			sort.Slice(helps, func(i, j int) bool {
				if helps[i].required != helps[j].required {
					return helps[i].required
				}
				return helps[i].name < helps[j].name
			})
			for _, pi := range helps {
				if pi.name == "owner" {
					var skip bool
					for _, hh := range helps {
						if hh.required && hh.name == "repo" {
							skip = true
							break
						}
					}
					if skip {
						continue
					}
				}
				nm := pi.name
				if pi.unsupported {
					nm = "~~" + nm + "~~"
				}
				desc := ""
				if pi.required {
					desc += "__Required__ "
				}
				if pi.unsupported {
					desc += "__unsupported by octo-cli__ "
				}
				desc += pi.description
				desc = strings.ReplaceAll(desc, "\n", "<br>")
				writerMust(buf.WriteString(fmt.Sprintf("| %s | %s |\n", nm, desc)))
			}
		}
	}

	return buf.Bytes(), nil
}

func removeHelpsWithName(helps []paramHelp, name string) []paramHelp {
	for {
		i := 0
		for ; i < len(helps); i++ {
			if helps[i].name == name {
				break
			}
		}
		if i == len(helps) {
			return helps
		}
		helps = append(helps[:i], helps[i+1:]...)
	}
}

func paramHelps(path, method string, op *openapi3.Operation, uops []unsupportedOptionalParam) []paramHelp {
	var result []paramHelp
	for i, parameter := range op.Parameters {
		if parameter.Value.Name == "accept" {
			continue
		}
		result = append(result, paramHelp{
			name:        parameter.Value.Name,
			required:    swaggerparser.ParamRequired(op.Parameters, i),
			description: parameter.Value.Description,
		})
	}

	previews, err := swaggerparser.OperationPreviews(op)
	if err != nil {
		panic(err)
	}

	for _, preview := range previews {
		note := preview.Note
		note = strings.Split(note, "```")[0]
		result = append(result, paramHelp{
			name:        preview.Name + "-preview",
			required:    preview.Required,
			description: note,
		})
	}

	var myUops []unsupportedOptionalParam
	for _, uop := range uops {
		if uop.routePath == path && uop.method == method {
			myUops = append(myUops, uop)
		}
	}
	for _, m := range swaggerparser.GetManualParamInfo(op) {
		result = removeHelpsWithName(result, m.Name)
		if m.Hidden {
			continue
		}
		result = append(result, paramHelp{
			name:        m.Name,
			required:    m.Required,
			description: m.Description,
		})
	}
	for _, bpi := range swaggerparser.GetBodyParamInfo(op, supported.RefFilter) {
		ph := paramHelp{
			name:        bpi.Name,
			required:    bpi.Required,
			description: bpi.Ref.Value.Description,
		}
		for _, uop := range myUops {
			if uop.paramName == bpi.Name {
				ph.unsupported = true
				break
			}
		}

		result = append(result, ph)
	}
	return result
}

func unsupportedHelp(swagger *openapi3.Swagger) ([]byte, error) {
	var buf bytes.Buffer

	svcData := map[string][]supported.UnsupportedOperation{}

	unsupOps := supported.GetUnsupportedOperations(swagger)
	for _, unsupOp := range unsupOps {
		opID := strings.Split(unsupOp.Operation.OperationID, "/")
		if len(opID) != 2 {
			return nil, fmt.Errorf("invalid OperationID")
		}
		svcName := util.ToArgName(opID[0])
		svcData[svcName] = append(svcData[svcName], unsupOp)
	}
	svcNames := make([]string, 0, len(svcData))
	for n := range svcData {
		svcNames = append(svcNames, n)
	}
	sort.Strings(svcNames)

	writerMust(writeMarkdownHeader(&buf, 1, "Unsupported Operations"))

	for _, svcName := range svcNames {
		for _, unsupOp := range svcData[svcName] {
			out := fmt.Sprintf(" - `%s %s` - %s\n",
				swaggerparser.GetOperationSvcName(unsupOp.Operation),
				swaggerparser.GetOperationName(unsupOp.Operation),
				unsupOp.Reason,
			)
			writerMust(
				buf.WriteString(out),
			)
		}
	}
	return buf.Bytes(), nil
}
