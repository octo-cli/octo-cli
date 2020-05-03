package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"github.com/spf13/afero"
	"golang.org/x/tools/imports"
)

var paramTypes = map[string]string{
	"string":    "string",
	"integer":   "int64",
	"string[]":  "[]string",
	"integer[]": "[]int64",
	"boolean":   "bool",
}

var updateMethodMap = map[string]string{
	"url":     "UpdateURLPath",
	"path":    "UpdateURLPath",
	"body":    "UpdateBody",
	"query":   "UpdateURLQuery",
	"header":  "AddRequestHeader",
	"preview": "UpdatePreview",
}

func dirFileMap(path string, fs afero.Fs) (map[string][]byte, error) {
	output := map[string][]byte{}
	fileInfos, err := afero.ReadDir(fs, path)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fileInfo.Sys()
		if fileInfo.IsDir() {
			continue
		}
		fileBytes, err := afero.ReadFile(fs, filepath.Join(path, fileInfo.Name()))
		if err != nil {
			return nil, err
		}
		output[fileInfo.Name()] = fileBytes
	}
	return output, nil
}

func verify(routesPath, outputPath string) ([]string, error) {
	realFs := afero.NewOsFs()
	tmpFs := afero.NewMemMapFs()
	tempDir, err := afero.TempDir(tmpFs, "", "")
	if err != nil {
		return nil, err
	}
	Generate(routesPath, tempDir, tmpFs)

	wantFiles, err := dirFileMap(tempDir, tmpFs)
	if err != nil {
		return nil, err
	}
	gotFiles, err := dirFileMap(outputPath, realFs)
	if err != nil {
		return nil, err
	}
	diffFiles := map[string]interface{}{}
	for name, wantBytes := range wantFiles {
		gotBytes := gotFiles[name]
		if !bytes.Equal(gotBytes, wantBytes) {
			diffFiles[name] = nil
		}
	}
	for name := range gotFiles {
		if _, ok := wantFiles[name]; !ok {
			diffFiles[name] = nil
		}
	}

	var output []string

	for v := range diffFiles {
		if strings.HasSuffix(v, "_test.go") ||
			strings.HasSuffix(v, "artisanally_handcrafted_code.go") {
			continue
		}
		output = append(output, v)
	}
	return output, nil
}

func Generate(routesPath, outputPath string, fs afero.Fs) {
	files, err := genFileTmpls(routesPath)
	if err != nil {
		panic(err)
	}
	for filename, fileTmpl := range files {
		err := writeGoFile(filename, "main", fileTmpl, outputPath, fs)
		if err != nil {
			panic(err)
		}
	}
}

func genFileTmpls(routesPath string) (map[string]FileTmpl, error) {
	CLITmpl := StructTmplHelper{
		Name: "CLI",
	}
	cmdHelps := map[string]map[string]string{}
	flagHelps := map[string]map[string]map[string]string{}
	svcTmpls := map[string]*SvcTmpl{}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(routesPath)
	if err != nil {
		return nil, err
	}
	for path, pathItem := range swagger.Paths {
		for method, op := range pathItem.Operations() {
			opID := strings.Split(op.OperationID, "/")
			if len(opID) != 2 {
				continue
			}
			idName := opID[1]
			svcNodeName := opID[0]
			svcName := strcase.ToCamel(opID[0])
			if cmdHelps[svcNodeName] == nil {
				cmdHelps[svcNodeName] = map[string]string{}
			}

			if flagHelps[svcNodeName] == nil {
				flagHelps[svcNodeName] = map[string]map[string]string{}
			}

			if svcTmpls[svcName] == nil {

				svcTmpls[svcName] = &SvcTmpl{
					SvcStruct: StructTmplHelper{
						Name: svcName + "Cmd",
					},
				}
				CLITmpl.Fields = append(CLITmpl.Fields, StructField{
					Name: svcName,
					Type: svcName + "Cmd",
					Tags: newTags(newTag("cmd", "")),
				})
			}

			structName := svcName + toArgName(idName) + "Cmd"
			tmplHelper := StructTmplHelper{
				Name:   structName,
				Fields: []StructField{{Type: "internal.BaseCmd"}},
			}

			runMethod := RunMethod{
				ReceiverName: structName,
				Method:       strings.ToUpper(method),
				URLPath:      path,
			}

			if flagHelps[svcNodeName][idName] == nil {
				flagHelps[svcNodeName][idName] = map[string]string{}
			}

			pd := propDataVals{
				fields:   &tmplHelper.Fields,
				rmParams: &runMethod.Params,
				helpers:  flagHelps[svcNodeName][idName],
			}

			err := previewData(op, pd)
			if err != nil {
				delete(flagHelps[svcNodeName], idName)
				continue
			}

			unsupportedBodyParams, err := bodyParamData(op, pd)
			if err != nil {
				fmt.Printf("%s has the following required params that are not supported:\n", op.OperationID)
				for s, b := range unsupportedBodyParams {
					if b {
						fmt.Println(s)
					}
				}
				delete(flagHelps[svcNodeName], idName)
				continue
			}

			err = paramData(op, pd)
			if err != nil {
				delete(flagHelps[svcNodeName], idName)
				continue
			}

			svcTmpl := svcTmpls[svcName]

			svcTmpl.CmdStructAndMethods = append(svcTmpl.CmdStructAndMethods, CmdStructAndMethod{
				CmdStruct: tmplHelper,
				RunMethod: runMethod,
			})

			helpText := op.Summary
			if op.ExternalDocs.URL != "" {
				helpText = fmt.Sprintf("%v - %v", helpText, op.ExternalDocs.URL)
			}
			cmdHelps[svcNodeName][idName] = helpText
			svcTmpl.SvcStruct.Fields = append(svcTmpl.SvcStruct.Fields, StructField{
				Name: toArgName(idName),
				Type: structName,
				Tags: newTags(newTag("cmd", "")),
			})
		}
	}
	sort.Slice(CLITmpl.Fields, func(i, j int) bool {
		return CLITmpl.Fields[i].Name < CLITmpl.Fields[j].Name
	})
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
		tmplSorting(svcTmpl)
		filename := strings.ToLower(svcTmpl.SvcStruct.Name) + ".go"
		files[filename] = FileTmpl{
			SvcTmpls: []SvcTmpl{*svcTmpl},
		}
	}
	return files, nil
}

func paramRequired(parameters openapi3.Parameters, idx int) bool {
	param := parameters[idx].Value
	if param.Name != "owner" || param.In != "path" || idx == len(parameters)-1 {
		return param.Required
	}
	nextParam := parameters[idx+1].Value
	return nextParam.Name != "repo" || nextParam.In != "path"
}

func previewData(op *openapi3.Operation, pd propDataVals) error {
	xMsg, ok := op.Extensions["x-github"].(json.RawMessage)
	if !ok {
		return nil
	}

	xg := struct {
		Legacy          bool
		EnabledForApps  bool
		GithubCloudOnly bool
		Previews        []struct {
			Name     string
			Required bool
			Note     string
		}
	}{}

	err := json.Unmarshal(xMsg, &xg)
	if err != nil {
		return err
	}
	for _, preview := range xg.Previews {
		tags := newTags(newTag("name", preview.Name+"-preview"))
		if preview.Required {
			setTag(tags, newTag("required", ""))
		}
		setTag(tags, newTag("help", preview.Note))
		*pd.fields = append(*pd.fields, StructField{
			Name: toArgName(preview.Name),
			Type: "bool",
			Tags: tags,
		})
		*pd.rmParams = append(*pd.rmParams, RunMethodParam{
			Name:         preview.Name,
			UpdateMethod: updateMethodMap["preview"],
			ValueField:   toArgName(preview.Name),
		})
	}
	return nil
}

func paramData(op *openapi3.Operation, pd propDataVals) error {
	for i, pRef := range op.Parameters {
		param := pRef.Value
		if param.Name == "accept" {
			continue
		}
		paramType, ok := paramTypes[param.Schema.Value.Type]
		if !ok {
			return fmt.Errorf("unexpected type %q for parameter %q", param.Schema.Value.Type, param.Name)
		}
		required := paramRequired(op.Parameters, i)
		tags := new(structtag.Tags)
		if required {
			setTag(tags, &structtag.Tag{Key: "required"})
		}
		setTag(tags, &structtag.Tag{Key: "name", Name: param.Name})
		*pd.fields = append(*pd.fields, StructField{
			Name: toArgName(param.Name),
			Type: paramType,
			Tags: tags,
		})
		*pd.rmParams = append(*pd.rmParams, RunMethodParam{
			Name:         param.Name,
			ValueField:   toArgName(param.Name),
			UpdateMethod: updateMethodMap[param.In],
		})
		pd.helpers[param.Name] = param.Description
	}
	return nil
}

type propDataVals struct {
	fields   *[]StructField
	rmParams *[]RunMethodParam
	helpers  map[string]string
}

func bodyParamData(op *openapi3.Operation, pd propDataVals) (map[string]bool, error) {
	if op.RequestBody == nil || op.RequestBody.Value.Content.Get("application/json") == nil {
		return nil, nil
	}
	bodySchema := op.RequestBody.Value.Content.Get("application/json").Schema.Value
	required := map[string]bool{}
	for _, s := range bodySchema.Required {
		required[s] = true
	}
	unsupported := map[string]bool{}
	for propertyName, prop := range bodySchema.Properties {
		singleBodyParam(propertyName, prop, unsupported, required, pd)
	}
	var err error
	for _, b := range unsupported {
		if b {
			err = fmt.Errorf("body has at least one required field with an unsupported type")
			break
		}
	}
	return unsupported, err
}

func singleBodyParam(propertyName string, prop *openapi3.SchemaRef, unsupported map[string]bool, required map[string]bool, pd propDataVals) {
	pv := prop.Value

	dotCount := strings.Count(propertyName, ".")
	if dotCount > 1 {
		fmt.Println(propertyName)
	}

	if pv.Type == "object" {
		if required[propertyName] {
			for _, s := range pv.Required {
				required[fmt.Sprintf("%s.%s", propertyName, s)] = true
			}
		}
		for subName, pRef := range pv.Properties {
			name := propertyName + "." + subName
			singleBodyParam(name, pRef, unsupported, required, pd)
		}
		return
	}

	paramType, ok := paramTypes[getPropType(pv)]
	if !ok {
		unsupported[propertyName] = required[propertyName]
		return
	}
	tags := new(structtag.Tags)
	if required[propertyName] {
		setTag(tags, &structtag.Tag{Key: "required"})
	}
	setTag(tags, &structtag.Tag{Key: "name", Name: propertyName})
	*pd.fields = append(*pd.fields, StructField{
		Name: toArgName(propertyName),
		Type: paramType,
		Tags: tags,
	})

	*pd.rmParams = append(*pd.rmParams, RunMethodParam{
		Name:         propertyName,
		UpdateMethod: updateMethodMap["body"],
		ValueField:   toArgName(propertyName),
	})
	pd.helpers[propertyName] = pv.Description
}

func sortCmdStructFields(fields []StructField) {
	if len(fields) == 0 {
		return
	}
	newFields := make([]StructField, 0, len(fields))
	holdouts := make([]StructField, 0, len(fields))
	for _, field := range fields {
		if field.Name == "" {
			holdouts = append(holdouts, field)
			continue
		}
		newFields = append(newFields, field)
	}
	sort.Slice(newFields, func(i, j int) bool {
		return newFields[i].Name < newFields[j].Name
	})
	sort.Slice(holdouts, func(i, j int) bool {
		return holdouts[i].Type < holdouts[j].Type
	})
	newFields = append(newFields, holdouts...)
	copy(fields, newFields)
}

func tmplSorting(svcTmpl *SvcTmpl) {
	sort.Slice(svcTmpl.SvcStruct.Fields, func(i, j int) bool {
		return svcTmpl.SvcStruct.Fields[i].Name < svcTmpl.SvcStruct.Fields[j].Name
	})
	for _, csm := range svcTmpl.CmdStructAndMethods {
		sortCmdStructFields(csm.CmdStruct.Fields)

		sort.Slice(csm.RunMethod.Params, func(i, j int) bool {
			return csm.RunMethod.Params[i].Name < csm.RunMethod.Params[j].Name
		})
	}
	sort.Slice(svcTmpl.CmdStructAndMethods, func(i, j int) bool {
		return svcTmpl.CmdStructAndMethods[i].CmdStruct.Name < svcTmpl.CmdStructAndMethods[j].CmdStruct.Name
	})
}

type RunMethodParam struct {
	Name         string
	ValueField   string
	UpdateMethod string
}

type RunMethod struct {
	ReceiverName string
	Method       string
	URLPath      string
	Params       []RunMethodParam
}

// StructField is one field in a StructTmplHelper
type StructField struct {
	Name string
	Type string
	Tags *structtag.Tags
}

type StructTmplHelper struct {
	Name   string
	Fields []StructField
}

type CmdStructAndMethod struct {
	CmdStruct StructTmplHelper
	RunMethod RunMethod
}

type SvcTmpl struct {
	SvcStruct           StructTmplHelper
	CmdStructAndMethods []CmdStructAndMethod
}

type FileTmpl struct {
	CmdHelps       map[string]map[string]string
	FlagHelps      map[string]map[string]map[string]string
	PrimaryStructs []StructTmplHelper
	SvcTmpls       []SvcTmpl
}

var tmpl = template.Must(template.New("").Parse(tmplt))

// language=GoTemplate
const tmplt = `
{{define "main"}}
// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

{{if .CmdHelps}}
var CmdHelps = map[string]map[string]string{
{{range $topCmd, $topCmdVals := .CmdHelps}}"{{$topCmd}}": {
{{range $cmd, $help := $topCmdVals}}"{{$cmd}}": {{printf "%q" $help}},
{{end}} 
},
{{end}} 
}
{{end}}

{{if .FlagHelps}}
var FlagHelps = map[string]map[string]map[string]string{
{{range $topCmd, $topCmdVals := .FlagHelps}}"{{$topCmd}}": {
{{range $cmd, $flagHelps := $topCmdVals}}"{{$cmd}}": {
{{range $flag, $help := $flagHelps}}"{{$flag}}": {{printf "%q" $help}},
{{end}}
}, 
{{end}}
}, 
{{end}}
} 
{{end}}


{{range .PrimaryStructs}}
{{template "structtype" .}}
{{end}}
{{range .SvcTmpls}}{{template "svctmpl" .}}{{end}}
{{end}}

{{define "svctmpl"}}{{if .SvcStruct}}{{template "structtype" .SvcStruct}}{{end}}
{{range .CmdStructAndMethods}}
{{if .CmdStruct}}{{template "structtype" .CmdStruct}}{{end}}
{{if .RunMethod}}{{template "runmethod" .RunMethod}}{{end}}
{{end}}
{{end}}

{{define "structtype"}}type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} {{if .Tags}}{{printf "%#q" .Tags}} {{end}}{{end}}
}{{end}}

{{define "runmethod"}}
func (c *{{.ReceiverName}}) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("{{.URLPath}}"){{range .Params}}{{template "runmethodparam" .}}{{end}}
	return c.DoRequest("{{.Method}}")
}{{end}}

{{define "runmethodparam"}}
	c.{{.UpdateMethod}}("{{.Name}}", c.{{.ValueField}}){{end}}
`

//newTag is a helper to create a new *structtag.Tag with fewer lines of code
func newTag(key, name string, options ...string) *structtag.Tag {
	return &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	}
}

//newTags creates a new *structtag.Tags from a list of tags
//  it will panic if one of the tags has no key
func newTags(tag ...*structtag.Tag) *structtag.Tags {
	tags := &structtag.Tags{}
	for _, tag := range tag {
		err := tags.Set(tag)
		if err != nil {
			panic(err)
		}
	}
	return tags
}

func setTag(tags *structtag.Tags, tag ...*structtag.Tag) {
	for _, tg := range tag {
		err := tags.Set(tg)
		if err != nil {
			panic(err)
		}
	}
}

//writeGoFile executes the named template and does the equivalent of `go fmt` and `goimports` on the output
func writeGoFile(filename, templateName string, p interface{}, path string, fs afero.Fs) error {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, templateName, p)
	if err != nil {
		return err
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(filename)
		fmt.Println(templateName)
		fmt.Println(buf.String())
		return err
	}
	out, err = imports.Process("", out, nil)
	if err != nil {
		return err
	}
	fl := filepath.Join(path, filename)
	return afero.WriteFile(fs, fl, out, 0644)
}

//toArgName takes input like "foo-bar" and returns "FooBar"
func toArgName(in string) string {
	out := in
	for _, separator := range []string{"_", "-", "."} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}

func getPropType(schema *openapi3.Schema) string {
	if schema == nil {
		return ""
	}
	if schema.Type != "array" || schema.Items == nil {
		return schema.Type
	}
	itemType := schema.Items.Value.Type
	return fmt.Sprintf("%s[]", itemType)
}
