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
	return writeFiles(files, outputPath, fs)
}

func cliTmpl(endpoints []model.Endpoint) structTmplHelper {
	result := structTmplHelper{
		name: "CLI",
	}
	for _, concern := range util.AllConcerns(endpoints) {
		result.fields = append(result.fields, structField{
			name:      util.ToArgName(concern),
			fieldType: util.ToArgName(concern) + "Cmd",
			tags:      map[string]string{"cmd": ""},
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
		if param.HelpText != "" {
			result[param.Name] = param.HelpText
		}
	}
	for _, param := range endpoint.QueryParams {
		if param.HelpText != "" {
			result[param.Name] = param.HelpText
		}
	}
	for _, param := range endpoint.Headers {
		if param.Name == "accept" {
			continue
		}
		if param.HelpText != "" {
			result[param.Name] = param.HelpText
		}
	}
	for _, preview := range endpoint.Previews {
		result[preview.Name+"-preview"] = util.FixPreviewNote(preview.Note)
	}
	if endpoint.JSONBodySchema != nil {
		bodyParams := util.FlattenParams(endpoint.JSONBodySchema.ObjectParams)
		for _, param := range bodyParams {
			if param.HelpText != "" {
				result[param.Name] = param.HelpText
			}
		}
	}
	mpi := overrides.GetManualParamInfo(endpoint.ID)
	for _, info := range mpi {
		delete(result, info.Name)
		if info.Description != "" {
			result[info.Name] = info.Description
		}
	}
	return result
}

func epFlagHelps(endpoints []model.Endpoint) map[string]map[string]map[string]string {
	result := map[string]map[string]map[string]string{}
	for _, endpoint := range endpoints {
		if util.EndpointIsUnsupported(endpoint) {
			continue
		}
		helps := endpointFieldHelp(endpoint)
		if len(helps) == 0 {
			continue
		}
		if result[endpoint.Concern] == nil {
			result[endpoint.Concern] = map[string]map[string]string{}
		}
		result[endpoint.Concern][endpoint.Name] = helps
	}
	return result
}

func epSvcTmpls(endpoints []model.Endpoint) map[string]*svcTmpl {
	result := map[string]*svcTmpl{}
	for _, concern := range util.AllConcerns(endpoints) {
		svcName := util.ToArgName(concern)
		result[svcName] = &svcTmpl{
			svcStruct: &structTmplHelper{
				name: svcName + "Cmd",
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
		result[svcName].cmdStructAndMethods = append(result[svcName].cmdStructAndMethods, cmdStructAndMethod{
			cmdStruct: cmdStruct,
			runMethod: &runMethod,
		})
		structName := svcName + util.ToArgName(endpoint.Name) + "Cmd"
		result[svcName].svcStruct.fields = append(result[svcName].svcStruct.fields, structField{
			name:      util.ToArgName(endpoint.Name),
			fieldType: structName,
			tags:      map[string]string{"cmd": ""},
		})
	}
	return result
}

func genFileTmpls(endpoints []model.Endpoint) map[string]fileTmpl {
	util.RemoveOwnerParams(endpoints)
	CLITmpl := cliTmpl(endpoints)
	cmdHelps := epCmdHelps(endpoints)
	flagHelps := epFlagHelps(endpoints)
	svcTmpls := epSvcTmpls(endpoints)
	files := map[string]fileTmpl{
		"cli.go": {
			cmdHelps:  cmdHelps,
			flagHelps: flagHelps,
			primaryStructs: []structTmplHelper{
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
		st := svcTmpls[key]
		filename := strings.ToLower(st.svcStruct.name) + ".go"
		files[filename] = fileTmpl{
			svcTmpls: []svcTmpl{*st},
		}
	}
	return files
}

func epPreviewCmdFields(endpoint model.Endpoint) []structField {
	result := make([]structField, 0, len(endpoint.Previews))
	for _, preview := range endpoint.Previews {
		result = append(result, structField{
			name:      util.ToArgName(preview.Name),
			fieldType: "bool",
			tags:      util.FieldTags(preview.Name+"-preview", preview.Required),
		})
	}
	return result
}

func manualCmdFields(opID string) []structField {
	mpi := overrides.GetManualParamInfo(opID)
	result := make([]structField, 0, len(mpi))
	for _, info := range mpi {
		tags := util.FieldTags(info.Name, info.Required)
		if info.Tags != nil {
			for k, v := range info.Tags {
				tags[k] = v
			}
		}
		result = append(result, structField{
			name:          util.ToArgName(info.Name),
			fieldType:     info.Type,
			tags:          tags,
			paramLocation: locBody,
		})
	}
	return result
}

func epBodyCmdFields(endpoint model.Endpoint) []structField {
	if endpoint.JSONBodySchema == nil {
		return nil
	}
	bodyParams := util.FlattenParams(endpoint.JSONBodySchema.ObjectParams)
	return paramsCmdFields(locBody, bodyParams)
}

func paramsCmdFields(loc paramLocation, params model.Params) []structField {
	result := make([]structField, 0, len(params))
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
		result = append(result, structField{
			name:          util.ToArgName(param.Name),
			fieldType:     paramType,
			tags:          util.FieldTags(param.Name, param.Required),
			paramLocation: loc,
			paramOrder:    paramOrder,
		})
	}
	return result
}

func endpointCmdStruct(endpoint model.Endpoint) *structTmplHelper {
	tmplHelper := structTmplHelper{
		name: endpointCmdStructName(endpoint),
		fields: []structField{
			{
				fieldType: "internal.BaseCmd",
			},
		},
	}
	tmplHelper.fields = append(tmplHelper.fields, epPreviewCmdFields(endpoint)...)
	tmplHelper.fields = append(tmplHelper.fields, epBodyCmdFields(endpoint)...)
	tmplHelper.fields = append(tmplHelper.fields, paramsCmdFields(locPath, endpoint.PathParams)...)
	tmplHelper.fields = append(tmplHelper.fields, paramsCmdFields(locQuery, endpoint.QueryParams)...)
	tmplHelper.fields = append(tmplHelper.fields, paramsCmdFields(locHeader, endpoint.Headers)...)
	mcf := manualCmdFields(endpoint.ID)
	for _, mField := range mcf {
		tmplHelper.fields = removeFieldsWithName(tmplHelper.fields, mField.name)
	}
	tmplHelper.fields = append(tmplHelper.fields, mcf...)
	return &tmplHelper
}

func removeFieldsWithName(fields []structField, name string) []structField {
	for {
		i := 0
		for ; i < len(fields); i++ {
			if fields[i].name == name {
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

func endpointRunMethod(endpoint model.Endpoint) runMethod {
	var cgas []codeGroupAdder
	for _, info := range overrides.GetManualParamInfo(endpoint.ID) {
		cgas = append(cgas, info.CodeAdder)
	}

	for _, header := range endpoint.Headers {
		if header.Name == "accept" {
			continue
		}
		cgas = append(cgas, newRunMethodAdder(header.Name, "c.AddRequestHeader"))
	}

	for _, param := range endpoint.PathParams {
		cgas = append(cgas, newRunMethodAdder(param.Name, "c.UpdateURLPath"))
	}

	for _, param := range endpoint.QueryParams {
		cgas = append(cgas, newRunMethodAdder(param.Name, "c.UpdateURLQuery"))
	}

	for _, preview := range endpoint.Previews {
		cgas = append(cgas, newRunMethodAdder(preview.Name, "c.UpdatePreview"))
	}

	if endpoint.JSONBodySchema != nil {
		for _, param := range util.FlattenParams(endpoint.JSONBodySchema.ObjectParams) {
			cgas = append(cgas, newRunMethodAdder(param.Name, "c.UpdateBody"))
		}
	}

	return runMethod{
		receiverName:    endpointCmdStructName(endpoint),
		method:          strings.ToUpper(endpoint.Method),
		urlPath:         endpoint.Path,
		codeGroupAdders: cgas,
	}
}
