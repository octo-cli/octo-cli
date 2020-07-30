package docgen

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/octo-cli/octo-cli/internal/generator/overrides"
	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/octo-cli/octo-cli/internal/model"
	"github.com/octo-cli/octo-cli/internal/model/openapi"
	"github.com/spf13/afero"
)

type paramHelp struct {
	name        string
	required    bool
	description string
	unsupported bool
}

type unsupportedOptionalParam struct {
	operationID string
	routePath   string
	method      string
	paramName   string
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
		if vals[i].operationID < vals[j].operationID {
			return true
		}
		if vals[i].operationID > vals[j].operationID {
			return false
		}
		return vals[i].paramName < vals[j].paramName
	})
}

func getUnsupportedOptionalParams2(endpoints []model.Endpoint) []unsupportedOptionalParam {
	var result []unsupportedOptionalParam
	for _, endpoint := range endpoints {
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		if endpoint.JSONBodySchema == nil {
			continue
		}
		for _, param := range endpoint.JSONBodySchema.ObjectParams {
			if util.IsSupportedModelParam(param) {
				continue
			}
			result = append(result, unsupportedOptionalParam{
				routePath:   endpoint.Path,
				method:      endpoint.Method,
				paramName:   param.Name,
				operationID: endpoint.ID,
			})
		}
	}
	sortUnsupportedOptionalParams(result)
	return result
}

func WriteDocs(schemaPath, docsPath string, fs afero.Fs) error {
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return err
	}
	endpoints, err := openapi.EndpointsFromSchema(schemaFile)
	if err != nil {
		return err
	}
	overrides.OverrideEndpoints(endpoints)
	util.RemoveOwnerParams(endpoints)
	opDoc, err := operationsHelp(endpoints)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, filepath.Join(docsPath, "operations.md"), opDoc, 0640)
	if err != nil {
		return err
	}
	return nil
}

func svcEndpoints(svc string, endpoints []model.Endpoint) []model.Endpoint {
	result := make([]model.Endpoint, 0, len(endpoints))
	for _, endpoint := range endpoints {
		if endpoint.Concern != svc {
			continue
		}
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		result = append(result, endpoint)
	}
	return result
}

func operationsHelp(endpoints []model.Endpoint) ([]byte, error) {
	var buf bytes.Buffer
	uops := getUnsupportedOptionalParams2(endpoints)
	for _, concern := range util.AllConcerns(endpoints) {
		eps := svcEndpoints(concern, endpoints)
		if len(eps) == 0 {
			continue
		}
		sort.Slice(eps, func(i, j int) bool {
			return eps[i].ID < eps[j].ID
		})
		writerMust(
			writeMarkdownHeader(&buf, 1, concern),
		)
		for _, opInfo := range eps {
			writerMust(writeMarkdownHeader(&buf, 2, fmt.Sprintf("%s %s", concern, opInfo.Name)))
			writerMust(buf.WriteString(opInfo.DocsURL + "\n\n"))
			writerMust(buf.WriteString(opInfo.HelpText + "\n"))
			helps := paramHelps(opInfo, uops)
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

func paramHelps(endpoint model.Endpoint, uops []unsupportedOptionalParam) []paramHelp {
	var result []paramHelp
	for _, param := range endpoint.PathParams {
		result = append(result, paramHelp{
			name:        param.Name,
			required:    param.Required,
			description: param.HelpText,
		})
	}
	for _, param := range endpoint.QueryParams {
		result = append(result, paramHelp{
			name:        param.Name,
			required:    param.Required,
			description: param.HelpText,
		})
	}
	for _, param := range endpoint.Headers {
		if param.Name == "accept" {
			continue
		}
		result = append(result, paramHelp{
			name:        param.Name,
			required:    param.Required,
			description: param.HelpText,
		})
	}
	for _, preview := range endpoint.Previews {
		note := util.FixPreviewNote(preview.Note)
		note = strings.Split(note, "```")[0]
		result = append(result, paramHelp{
			name:        preview.Name + "-preview",
			required:    preview.Required,
			description: note,
		})
	}
	var myUops []unsupportedOptionalParam
	for _, uop := range uops {
		if uop.routePath == endpoint.Path && uop.method == endpoint.Method {
			myUops = append(myUops, uop)
		}
	}
	for _, m := range overrides.GetManualParamInfo(endpoint.ID) {
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
	if endpoint.JSONBodySchema == nil {
		return result
	}
	bodyParams := util.FlattenParams(endpoint.JSONBodySchema.ObjectParams)
	for _, param := range bodyParams {
		ph := paramHelp{
			name:        param.Name,
			required:    param.Required,
			description: param.HelpText,
		}
		for _, uop := range myUops {
			if uop.paramName == param.Name {
				ph.unsupported = true
				break
			}
		}
		result = append(result, ph)
	}
	return result
}
