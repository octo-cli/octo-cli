package codegen

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/octo-cli/octo-cli/internal/generator/supported"
	"github.com/octo-cli/octo-cli/internal/generator/swaggerparser"
	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/spf13/afero"
)

func Generate(routesPath, outputPath string, fs afero.Fs) error {
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(routesPath)
	if err != nil {
		return err
	}
	files, err := GenFileTmpls(swagger)
	if err != nil {
		return err
	}
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

func cliTmpl(swagger *openapi3.Swagger) StructTmplHelper {
	result := StructTmplHelper{
		Name: "CLI",
	}
	for _, svcName := range swaggerparser.AllServiceNames(swagger) {
		if len(supported.SvcOperations(swagger, svcName)) == 0 {
			continue
		}
		result.Fields = append(result.Fields, StructField{
			Name: util.ToArgName(svcName),
			Type: util.ToArgName(svcName) + "Cmd",
			Tags: util.NewTags(util.NewTag("cmd", "")),
		})
	}
	return result
}

func operationCmdHelp(op *openapi3.Operation) string {
	if op.ExternalDocs.URL == "" {
		return op.Summary
	}
	return fmt.Sprintf("%v - %v", op.Summary, op.ExternalDocs.URL)
}

func swCmdHelps(swagger *openapi3.Swagger) map[string]map[string]string {
	result := map[string]map[string]string{}
	swaggerparser.ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		if supported.OperationIsUnsupported(op, path, method) {
			return
		}
		idName := swaggerparser.GetOperationName(op)
		svcNodeName := swaggerparser.GetOperationSvcName(op)
		if result[svcNodeName] == nil {
			result[svcNodeName] = map[string]string{}
		}
		result[svcNodeName][idName] = operationCmdHelp(op)
	})
	return result
}

func manualFieldHelp(op *openapi3.Operation) map[string]string {
	mpi := swaggerparser.GetManualParamInfo(op)
	result := make(map[string]string, len(mpi))
	for _, info := range mpi {
		result[info.Name] = info.Description
	}
	return result
}

func bodyFieldHelp(op *openapi3.Operation) map[string]string {
	bpi := swaggerparser.GetBodyParamInfo(op, supported.RefFilter)
	result := make(map[string]string, len(bpi))
	for _, info := range bpi {
		if !supported.IsSupportedParam(info.Ref) {
			continue
		}
		result[info.Name] = info.Ref.Value.Description
	}
	return result
}

func paramFieldHelp(op *openapi3.Operation) map[string]string {
	result := make(map[string]string, len(op.Parameters))
	for _, pRef := range op.Parameters {
		param := pRef.Value
		if param.Name == "accept" {
			continue
		}
		result[param.Name] = param.Description
	}
	return result
}

func previewFieldHelp(op *openapi3.Operation) map[string]string {
	previews, err := swaggerparser.OperationPreviews(op)
	if err != nil {
		return nil
	}
	result := make(map[string]string, len(previews))
	for _, preview := range previews {
		result[preview.Name+"-preview"] = preview.Note
	}
	return result
}

func operationFieldHelp(op *openapi3.Operation) map[string]string {
	result := bodyFieldHelp(op)
	for k, v := range paramFieldHelp(op) {
		result[k] = v
	}
	for k, v := range previewFieldHelp(op) {
		result[k] = v
	}
	for k, v := range manualFieldHelp(op) {
		result[k] = v
	}
	return result
}

func swFlagHelps(swagger *openapi3.Swagger) map[string]map[string]map[string]string {
	result := map[string]map[string]map[string]string{}
	swaggerparser.ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		if supported.OperationIsUnsupported(op, path, method) {
			return
		}
		idName := swaggerparser.GetOperationName(op)
		svcNodeName := swaggerparser.GetOperationSvcName(op)
		if result[svcNodeName] == nil {
			result[svcNodeName] = map[string]map[string]string{}
		}
		result[svcNodeName][idName] = operationFieldHelp(op)
	})
	return result
}

func swSvcTmpls(swagger *openapi3.Swagger) (map[string]*SvcTmpl, error) {
	result := map[string]*SvcTmpl{}
	err := swaggerparser.ForEachOperationErr(swagger, func(path, method string, op *openapi3.Operation) error {
		if supported.OperationIsUnsupported(op, path, method) {
			return nil
		}
		idName := swaggerparser.GetOperationName(op)
		svcName := util.ToArgName(swaggerparser.GetOperationSvcName(op))
		if result[svcName] == nil {
			result[svcName] = &SvcTmpl{
				SvcStruct: StructTmplHelper{
					Name: svcName + "Cmd",
				},
			}
		}

		cmdStruct, err := operationCmdStruct(op)
		if err != nil {
			return err
		}
		runMethod, err := operationRunMethod(op, method, path)
		if err != nil {
			return err
		}
		result[svcName].CmdStructAndMethods = append(result[svcName].CmdStructAndMethods, CmdStructAndMethod{
			CmdStruct: *cmdStruct,
			RunMethod: runMethod,
		})
		structName := svcName + util.ToArgName(idName) + "Cmd"
		result[svcName].SvcStruct.Fields = append(result[svcName].SvcStruct.Fields, StructField{
			Name: util.ToArgName(idName),
			Type: structName,
			Tags: util.NewTags(util.NewTag("cmd", "")),
		})
		return nil
	})
	return result, err
}

func GenFileTmpls(swagger *openapi3.Swagger) (map[string]FileTmpl, error) {
	err := swaggerparser.RemoveOwnerParams(swagger)
	if err != nil {
		return nil, err
	}
	CLITmpl := cliTmpl(swagger)
	cmdHelps := swCmdHelps(swagger)
	flagHelps := swFlagHelps(swagger)
	svcTmpls, err := swSvcTmpls(swagger)
	if err != nil {
		return nil, err
	}
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
	return files, nil
}

func previewCmdFields(op *openapi3.Operation) ([]StructField, error) {
	previews, err := swaggerparser.OperationPreviews(op)
	if err != nil {
		return nil, err
	}
	result := make([]StructField, 0, len(previews))
	for _, preview := range previews {
		result = append(result, StructField{
			Name: util.ToArgName(preview.Name),
			Type: "bool",
			Tags: util.FieldTags(preview.Name+"-preview", preview.Required),
		})
	}
	return result, nil
}

func bodyCmdFields(op *openapi3.Operation) []StructField {
	if op.RequestBody == nil {
		return nil
	}
	pis := swaggerparser.GetBodyParamInfo(op, supported.RefFilter)
	result := make([]StructField, 0, len(pis))
	for _, pi := range pis {
		tp, ok := util.ParamTypes[util.GetPropType(pi.Ref.Value)]
		if !ok {
			continue
		}
		result = append(result, StructField{
			Name:          util.ToArgName(pi.Name),
			Type:          tp,
			Tags:          util.FieldTags(pi.Name, pi.Required),
			ParamLocation: locBody,
		})
	}
	return result
}

func manualCmdFields(op *openapi3.Operation) ([]StructField, error) {
	mpi := swaggerparser.GetManualParamInfo(op)
	result := make([]StructField, 0, len(mpi))
	for _, info := range mpi {
		tags := util.FieldTags(info.Name, info.Required)
		if info.Tags != nil {
			for _, infoTag := range info.Tags.Tags() {
				err := tags.Set(infoTag)
				if err != nil {
					return nil, err
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
	return result, nil
}

func paramCmdFields(op *openapi3.Operation) []StructField {
	result := make([]StructField, 0, len(op.Parameters))
	for i, pRef := range op.Parameters {
		param := pRef.Value
		if param.Name == "accept" {
			continue
		}
		paramType, ok := util.ParamTypes[param.Schema.Value.Type]
		if !ok {
			continue
		}
		required := swaggerparser.ParamRequired(op.Parameters, i)
		var loc paramLocation
		switch param.In {
		case openapi3.ParameterInPath:
			loc = locPath
		case openapi3.ParameterInHeader:
			loc = locHeader
		case openapi3.ParameterInQuery:
			loc = locQuery
		}
		var paramOrder int
		if loc == locPath {
			paramOrder = i
		}
		result = append(result, StructField{
			Name:          util.ToArgName(param.Name),
			Type:          paramType,
			Tags:          util.FieldTags(param.Name, required),
			ParamLocation: loc,
			ParamOrder:    paramOrder,
		})
	}
	return result
}

func operationCmdStruct(op *openapi3.Operation) (*StructTmplHelper, error) {
	tmplHelper := StructTmplHelper{
		Name: operationCmdStructName(op),
		Fields: []StructField{
			{
				Type:   "internal.BaseCmd",
				Import: "github.com/octo-cli/octo-cli/internal",
			},
		},
	}
	previewFields, err := previewCmdFields(op)
	if err != nil {
		return nil, err
	}
	tmplHelper.Fields = append(tmplHelper.Fields, previewFields...)
	tmplHelper.Fields = append(tmplHelper.Fields, bodyCmdFields(op)...)
	tmplHelper.Fields = append(tmplHelper.Fields, paramCmdFields(op)...)
	mcf, err := manualCmdFields(op)
	if err != nil {
		return nil, err
	}
	for _, mField := range mcf {
		tmplHelper.Fields = removeFieldsWithName(tmplHelper.Fields, mField.Name)
	}
	tmplHelper.Fields = append(tmplHelper.Fields, mcf...)
	return &tmplHelper, nil
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

func operationCmdStructName(op *openapi3.Operation) string {
	svcName := util.ToArgName(swaggerparser.GetOperationSvcName(op))
	idName := swaggerparser.GetOperationName(op)
	return svcName + util.ToArgName(idName) + "Cmd"
}

func bodyCodeBlocks(op *openapi3.Operation) ([]CodeBlock, error) {
	if op.RequestBody == nil {
		return nil, nil
	}
	pis := swaggerparser.GetBodyParamInfo(op, supported.RefFilter)
	result := make([]CodeBlock, 0, len(pis))
	var buf bytes.Buffer
	for _, pi := range pis {
		if !supported.IsSupportedParam(pi.Ref) {
			continue
		}
		buf.Reset()
		err := tmpl.ExecuteTemplate(&buf, "RunMethodParam", RunMethodParam{
			Name:         pi.Name,
			ValueField:   util.ToArgName(pi.Name),
			UpdateMethod: UpdateMethodMap["body"],
		})
		if err != nil {
			return nil, err
		}
		result = append(result, CodeBlock{
			Code: buf.String(),
		})
	}
	return result, nil
}

func manualCodeBlocks(op *openapi3.Operation) []CodeBlock {
	mpi := swaggerparser.GetManualParamInfo(op)
	result := make([]CodeBlock, 0, len(mpi))
	for _, info := range mpi {
		result = append(result, CodeBlock{
			Code:    info.RunCode,
			Imports: info.CodeImports,
		})
	}
	return result
}

func pathCodeBlocks(op *openapi3.Operation) ([]CodeBlock, error) {
	result := make([]CodeBlock, 0, len(op.Parameters))
	var buf bytes.Buffer
	for _, pRef := range op.Parameters {
		param := pRef.Value
		if param.Name == "accept" {
			continue
		}
		buf.Reset()
		err := tmpl.ExecuteTemplate(&buf, "RunMethodParam", RunMethodParam{
			Name:         param.Name,
			ValueField:   util.ToArgName(param.Name),
			UpdateMethod: UpdateMethodMap[param.In],
		})
		if err != nil {
			return nil, err
		}
		result = append(result, CodeBlock{
			Code: buf.String(),
		})
	}
	return result, nil
}

func previewCodeBlocks(op *openapi3.Operation) ([]CodeBlock, error) {
	previews, err := swaggerparser.OperationPreviews(op)
	if err != nil {
		return nil, nil
	}
	result := make([]CodeBlock, 0, len(previews))
	var buf bytes.Buffer
	for _, preview := range previews {
		buf.Reset()
		err := tmpl.ExecuteTemplate(&buf, "RunMethodParam", RunMethodParam{
			Name:         preview.Name,
			UpdateMethod: UpdateMethodMap["preview"],
			ValueField:   util.ToArgName(preview.Name),
		})
		if err != nil {
			return nil, err
		}
		result = append(result, CodeBlock{
			Code: buf.String(),
		})
	}
	return result, nil
}

func operationRunMethod(op *openapi3.Operation, method, path string) (RunMethod, error) {
	runMethod := RunMethod{
		ReceiverName: operationCmdStructName(op),
		Method:       strings.ToUpper(method),
		URLPath:      path,
	}

	runMethod.CodeBlocks = manualCodeBlocks(op)

	pathBlocks, err := pathCodeBlocks(op)
	if err != nil {
		return runMethod, err
	}
	runMethod.CodeBlocks = append(runMethod.CodeBlocks, pathBlocks...)

	previewBlocks, err := previewCodeBlocks(op)
	if err != nil {
		return runMethod, err
	}
	runMethod.CodeBlocks = append(runMethod.CodeBlocks, previewBlocks...)

	bodyBlocks, err := bodyCodeBlocks(op)
	if err != nil {
		return runMethod, err
	}
	runMethod.CodeBlocks = append(runMethod.CodeBlocks, bodyBlocks...)

	return runMethod, nil
}
