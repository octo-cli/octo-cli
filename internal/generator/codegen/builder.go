package codegen

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/octo-cli/octo-cli/internal/generator/overrides"
	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/octo-cli/octo-cli/internal/model"
	"github.com/octo-cli/octo-cli/internal/model/openapi"
	"github.com/spf13/afero"
)

func Generate(routesPath, outputPath string, fs afero.Fs) error {
	schemaFile, err := os.Open(routesPath)
	if err != nil {
		return err
	}
	endpoints, err := openapi.EndpointsFromSchema(schemaFile)
	if err != nil {
		return err
	}
	files := genFileTmpls(endpoints)
	return WriteFiles(files, outputPath, fs)
}

var UpdateMethodMap = map[string]string{
	"url":     "c.UpdateURLPath",
	"path":    "c.UpdateURLPath",
	"body":    "c.UpdateBody",
	"query":   "c.UpdateURLQuery",
	"header":  "c.AddRequestHeader",
	"preview": "c.UpdatePreview",
}

func cliTmpl(endpoints []model.Endpoint) StructTmplHelper {
	result := StructTmplHelper{
		Name: "CLI",
	}
	for _, concern := range util.AllConcerns(endpoints) {
		result.Fields = append(result.Fields, StructField{
			Name: util.ToArgName(concern),
			Type: util.ToArgName(concern) + "Cmd",
			Tags: util.NewTags(util.NewTag("cmd", "")),
		})
	}
	return result
}

func endpointCmdHelp(endpoint model.Endpoint) string {
	if endpoint.DocsURL == "" {
		return endpoint.Summary
	}
	return fmt.Sprintf("%v - %v", endpoint.Summary, endpoint.DocsURL)
}

func epCmdHelps(endpoints []model.Endpoint) map[string]map[string]string {
	result := map[string]map[string]string{}
	for _, concern := range util.AllConcerns(endpoints) {
		result[concern] = map[string]string{}
	}
	for _, endpoint := range endpoints {
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		result[endpoint.Concern][endpoint.Name] = endpointCmdHelp(endpoint)
	}
	return result
}

func endpointFieldHelp(endpoint model.Endpoint) map[string]string {
	result := map[string]string{}
	for _, param := range endpoint.PathParams {
		result[param.Name] = param.HelpText
	}
	for _, param := range endpoint.QueryParams {
		result[param.Name] = param.HelpText
	}
	for _, param := range endpoint.Headers {
		if param.Name == "accept" {
			continue
		}
		result[param.Name] = param.HelpText
	}
	for _, preview := range endpoint.Previews {
		result[preview.Name+"-preview"] = util.FixPreviewNote(preview.Note)
	}
	if endpoint.JSONBodySchema != nil {
		bodyParams := util.FlattenParams(endpoint.JSONBodySchema.ObjectParams)
		for _, param := range bodyParams {
			result[param.Name] = param.HelpText
		}
	}
	mpi := overrides.GetManualParamInfo(endpoint.ID)
	for _, info := range mpi {
		result[info.Name] = info.Description
	}
	return result
}

func epFlagHelps(endpoints []model.Endpoint) map[string]map[string]map[string]string {
	result := map[string]map[string]map[string]string{}
	for _, concern := range util.AllConcerns(endpoints) {
		result[concern] = map[string]map[string]string{}
	}
	for _, endpoint := range endpoints {
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		result[endpoint.Concern][endpoint.Name] = endpointFieldHelp(endpoint)
	}
	return result
}

func epSvcTmpls(endpoints []model.Endpoint) map[string]*SvcTmpl {
	result := map[string]*SvcTmpl{}
	for _, concern := range util.AllConcerns(endpoints) {
		svcName := util.ToArgName(concern)
		result[svcName] = &SvcTmpl{
			SvcStruct: StructTmplHelper{
				Name: svcName + "Cmd",
			},
		}
	}
	for _, endpoint := range endpoints {
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		svcName := util.ToArgName(endpoint.Concern)
		cmdStruct := endpointCmdStruct(endpoint)
		runMethod := endpointRunMethod(endpoint)
		result[svcName].CmdStructAndMethods = append(result[svcName].CmdStructAndMethods, CmdStructAndMethod{
			CmdStruct: *cmdStruct,
			RunMethod: runMethod,
		})
		structName := svcName + util.ToArgName(endpoint.Name) + "Cmd"
		result[svcName].SvcStruct.Fields = append(result[svcName].SvcStruct.Fields, StructField{
			Name: util.ToArgName(endpoint.Name),
			Type: structName,
			Tags: util.NewTags(util.NewTag("cmd", "")),
		})
	}
	return result
}

func genFileTmpls(endpoints []model.Endpoint) map[string]FileTmpl {
	util.RemoveOwnerParams(endpoints)
	CLITmpl := cliTmpl(endpoints)
	cmdHelps := epCmdHelps(endpoints)
	flagHelps := epFlagHelps(endpoints)
	svcTmpls := epSvcTmpls(endpoints)
	files := map[string]FileTmpl{
		"cli.go": {
			CmdHelps:  cmdHelps,
			FlagHelps: flagHelps,
			PrimaryStructs: []StructTmplHelper{
				CLITmpl,
			},
		},
	}
	svcTmplsKeys := make([]string, 0, len(svcTmpls))
	for k := range svcTmpls {
		svcTmplsKeys = append(svcTmplsKeys, k)
	}
	sort.Strings(svcTmplsKeys)
	for _, key := range svcTmplsKeys {
		svcTmpl := svcTmpls[key]
		filename := strings.ToLower(svcTmpl.SvcStruct.Name) + ".go"
		files[filename] = FileTmpl{
			SvcTmpls: []SvcTmpl{*svcTmpl},
		}
	}
	return files
}

func epPreviewCmdFields(endpoint model.Endpoint) []StructField {
	result := make([]StructField, 0, len(endpoint.Previews))
	for _, preview := range endpoint.Previews {
		result = append(result, StructField{
			Name: util.ToArgName(preview.Name),
			Type: "bool",
			Tags: util.FieldTags(preview.Name+"-preview", preview.Required),
		})
	}
	return result
}

func manualCmdFields(opID string) []StructField {
	mpi := overrides.GetManualParamInfo(opID)
	result := make([]StructField, 0, len(mpi))
	for _, info := range mpi {
		tags := util.FieldTags(info.Name, info.Required)
		if info.Tags != nil {
			for _, infoTag := range info.Tags.Tags() {
				err := tags.Set(infoTag)
				if err != nil {
					panic(err)
				}
			}
		}
		result = append(result, StructField{
			Name:          util.ToArgName(info.Name),
			Type:          info.Type,
			Tags:          tags,
			Import:        info.FieldImport,
			ParamLocation: locBody,
		})
	}
	return result
}

func epBodyCmdFields(endpoint model.Endpoint) []StructField {
	if endpoint.JSONBodySchema == nil {
		return nil
	}
	bodyParams := util.FlattenParams(endpoint.JSONBodySchema.ObjectParams)
	return paramsCmdFields(locBody, bodyParams)
}

func paramsCmdFields(loc paramLocation, params model.Params) []StructField {
	result := make([]StructField, 0, len(params))
	for i, param := range params {
		if param.Name == "accept" {
			continue
		}
		paramType := util.SchemaParamType(param.Schema)
		if paramType == "" || paramType == "[]" {
			continue
		}
		var paramOrder int
		if loc == locPath {
			paramOrder = i
		}
		result = append(result, StructField{
			Name:          util.ToArgName(param.Name),
			Type:          paramType,
			Tags:          util.FieldTags(param.Name, param.Required),
			ParamLocation: loc,
			ParamOrder:    paramOrder,
		})
	}
	return result
}

func endpointCmdStruct(endpoint model.Endpoint) *StructTmplHelper {
	tmplHelper := StructTmplHelper{
		Name: endpointCmdStructName(endpoint),
		Fields: []StructField{
			{
				Type:   "internal.BaseCmd",
				Import: "github.com/octo-cli/octo-cli/internal",
			},
		},
	}
	tmplHelper.Fields = append(tmplHelper.Fields, epPreviewCmdFields(endpoint)...)
	tmplHelper.Fields = append(tmplHelper.Fields, epBodyCmdFields(endpoint)...)
	tmplHelper.Fields = append(tmplHelper.Fields, paramsCmdFields(locPath, endpoint.PathParams)...)
	tmplHelper.Fields = append(tmplHelper.Fields, paramsCmdFields(locQuery, endpoint.QueryParams)...)
	tmplHelper.Fields = append(tmplHelper.Fields, paramsCmdFields(locHeader, endpoint.Headers)...)
	mcf := manualCmdFields(endpoint.ID)
	for _, mField := range mcf {
		tmplHelper.Fields = removeFieldsWithName(tmplHelper.Fields, mField.Name)
	}
	tmplHelper.Fields = append(tmplHelper.Fields, mcf...)
	return &tmplHelper
}

func removeFieldsWithName(fields []StructField, name string) []StructField {
	for {
		i := 0
		for ; i < len(fields); i++ {
			if fields[i].Name == name {
				break
			}
		}
		if i == len(fields) {
			return fields
		}
		fields = append(fields[:i], fields[i+1:]...)
	}
}

func endpointCmdStructName(endpoint model.Endpoint) string {
	return util.ToArgName(endpoint.Concern) + util.ToArgName(endpoint.Name) + "Cmd"
}

func endpointRunMethod(endpoint model.Endpoint) RunMethod {
	var blocks []CodeBlock
	for _, info := range overrides.GetManualParamInfo(endpoint.ID) {
		blocks = append(blocks, CodeBlock{
			Code:    info.RunCode,
			Imports: info.CodeImports,
		})
	}

	for _, header := range endpoint.Headers {
		if header.Name == "accept" {
			continue
		}
		blocks = append(blocks, newRunMethodCodeBlock(header.Name, "c.AddRequestHeader"))
	}

	for _, param := range endpoint.PathParams {
		blocks = append(blocks, newRunMethodCodeBlock(param.Name, "c.UpdateURLPath"))
	}

	for _, param := range endpoint.QueryParams {
		blocks = append(blocks, newRunMethodCodeBlock(param.Name, "c.UpdateURLQuery"))
	}

	for _, preview := range endpoint.Previews {
		blocks = append(blocks, newRunMethodCodeBlock(preview.Name, UpdateMethodMap["preview"]))
	}

	if endpoint.JSONBodySchema != nil {
		for _, param := range util.FlattenParams(endpoint.JSONBodySchema.ObjectParams) {
			blocks = append(blocks, newRunMethodCodeBlock(param.Name, "c.UpdateBody"))
		}
	}

	return RunMethod{
		ReceiverName: endpointCmdStructName(endpoint),
		Method:       strings.ToUpper(endpoint.Method),
		URLPath:      endpoint.Path,
		CodeBlocks:   blocks,
	}
}
