package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/alecthomas/kong"
	"github.com/fatih/camelcase"
	"github.com/fatih/structtag"
	"github.com/google/go-github/github"
	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

type (
	svc struct {
		Name     string
		Commands []cmd
	}

	pkg struct {
		PackageName   string
		Imports       []string
		CmdHelpers    []*structTmplHelper
		OptionStructs []optionsStruct
	}

	cmd struct {
		Name     string
		ArgNames []string
		Route    *svcRoute
	}

	routeParams []*routeParam

	routeParam struct {
		Name        string
		Type        string
		Description string
		Default     interface{}
		Required    bool
		Enum        []interface{}
		Location    string
		MapTo       string
	}

	svcRoutes []*svcRoute

	rtServices map[string]svcRoutes

	valSetter struct {
		TargetIsPtr bool
		Name        string
		FlagName    string
	}

	toFuncTmplHelper struct {
		ReceiverName         string
		TargetName           string
		TargetType           string
		ValSetters           []valSetter
		IncludePointerHelper bool
	}

	svcRoute struct {
		Description      string
		Method           string
		Path             string
		Name             string
		EnabledForApps   bool
		IDName           string
		DocumentationURL string
		Params           routeParams
		Requests []interface{}
		Responses []struct {
			Headers struct {
				Status string
				ContentType string `json:"content-type"`
			}
			Body interface{}
			Description interface{}
		}
	}

	runMethodArgHelper struct {
		Name  string
		IsPtr bool
	}

	runMethodHelper struct {
		StructName string
		HasElement bool
		SvcName    string
		FuncName   string
		Args       []runMethodArgHelper
	}

	structTmplHelper struct {
		Name           string
		Fields         []structField
		RunMethod      *runMethodHelper
		OptionsStructs []optionsStruct
	}

	structField struct {
		Name string
		Type string
		Tags *structtag.Tags
	}

	optionsStruct struct {
		StructName string
		MainStruct structTmplHelper
		ToFunc     toFuncTmplHelper
	}

	configHcl struct {
		Service map[string]struct {
			RouteName string
			Command   map[string]struct {
				RoutesName string
				ArgNames   []string
			}
		}
	}

	genCli struct {
		Run            genCliRun            `cmd:"" help:"generate production code"`
		UpdateRoutes   genCliUpdateRoutes   `cmd:"" help:"update routes.json with the latest"`
		UpdateTestdata genCliUpdateTestdata `cmd:"" help:"updates routes.json and exampleapp in generator/testdata"`
	}

	genCliUpdateRoutes struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	genCliRun struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		ConfigFile string `type:"existingfile" default:"config.hcl"`
		OutputPath string `type:"existingdir" default:"."`
	}

	genCliUpdateTestdata struct{}
)

var (
	pkgImports = []string{
		"context",
		"encoding/json",
		"github.com/alecthomas/kong",
		"github.com/google/go-github/github",
		"golang.org/x/oauth2",
		"time",
	}

	tmpl = template.Must(template.New("").Parse(pkgTemplate))
)

func (k *genCliRun) Run() error {
	svcs, err := buildSvcs(k.RoutesPath, k.ConfigFile)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = writeOutput(k.OutputPath, svcs)
	return errors.Wrap(err, "")
}

//noinspection GoUnhandledErrorResult
func (k *genCliUpdateRoutes) Run() error {
	resp, err := http.Get(k.RoutesURL)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer resp.Body.Close()
	outFile, err := os.Create(k.RoutesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	return errors.Wrap(err, "")
}

//noinspection GoUnhandledErrorResult
func (k *genCliUpdateTestdata) Run() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "generator/testdata/routes.json"
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer resp.Body.Close()
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "")
	}
	svcs, err := buildSvcs(routesPath, "generator/testdata/exampleapp_config.hcl")
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = writeOutput("generator/testdata/exampleapp", svcs)
	return errors.Wrap(err, "")
}

func buildSvcs(routesPath, configFile string) ([]svc, error) {
	rt, err := getRtServices(routesPath)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	hConfig := &configHcl{}
	hConfigBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	err = hcl.Decode(hConfig, string(hConfigBytes))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var svcs []svc

	for svcName, hsvc := range hConfig.Service {
		svcRoutesName := hsvc.RouteName
		if svcRoutesName == "" {
			svcRoutesName = flagName(svcName)
		}
		var cmdNames []string
		for cmdName := range hsvc.Command {
			cmdNames = append(cmdNames, cmdName)
		}
		sort.Strings(cmdNames)
		cmds := make([]cmd, len(cmdNames))
		for i, cmdName := range cmdNames {
			hcmd := hsvc.Command[cmdName]
			routesName := hcmd.RoutesName
			if routesName == "" {
				routesName = flagName(cmdName)
			}
			cmds[i] = newCmd(cmdName, rt.svcRoutes(svcRoutesName).findByIdName(routesName), hcmd.ArgNames...)
		}
		svcs = append(svcs, svc{
			Name:     svcName,
			Commands: cmds,
		})
	}
	return svcs, nil
}

func main() {
	k := kong.Parse(&genCli{})
	err := k.Run()
	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println(err)
	}
	k.FatalIfErrorf(err)
}

func writeOutput(outputPath string, svcs []svc) error {
	for _, s := range svcs {
		p, err := s.pkg()
		if err != nil {
			return errors.Wrap(err, "")
		}
		err = p.writeToDir(filepath.Join(outputPath, p.outputDir("services")))
		if err != nil {
			return errors.Wrap(err, "")
		}
	}
	return nil
}

func newCmd(name string, rt *svcRoute, argNames ...string) cmd {
	if len(argNames) == 0 {
		for _, p := range rt.Params {
			if p.Location == "url" {
				argNames = append(argNames, toArgName(p.Name))
			}
		}
	}

	return cmd{
		Name:     name,
		ArgNames: argNames,
		Route:    rt,
	}
}

func (c *cmd) tags() (*structtag.Tags, error) {
	tags := &structtag.Tags{}
	err := tags.Set(newTag("cmd", ""))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	err = tags.Set(&structtag.Tag{Key: "help", Name: c.Route.Name})
	return tags, errors.Wrap(err, "")
}

func newStructField(name, ftype string, tgs *structtag.Tags) *structField {
	return &structField{
		Name: name,
		Type: ftype,
		Tags: tgs,
	}
}

func newTag(key, name string, options ...string) *structtag.Tag {
	return &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	}
}

func newTags(tags ...*structtag.Tag) (*structtag.Tags, error) {
	tgs := &structtag.Tags{}
	for _, tag := range tags {
		err := tgs.Set(tag)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
	}
	return tgs, nil
}

func newSvcCmd(svcName string, cmds []cmd) (*structTmplHelper, error) {
	var fields []structField
	for _, cmd := range cmds {
		tgs, err := cmd.tags()
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		fields = append(fields, structField{
			Name: cmd.Name,
			Type: svcName + cmd.Name + "Cmd",
			Tags: tgs,
		})
	}
	return &structTmplHelper{
		Name:   svcName + "Cmd",
		Fields: fields,
	}, nil
}

func (s *svc) getStructField() (reflect.StructField, bool) {
	clientType := reflect.TypeOf(github.Client{})
	return clientType.FieldByName(s.Name)
}

func (s *svc) pkg() (*pkg, error) {
	field, ok := s.getStructField()
	if !ok {
		return nil, errors.New("can't find structField")
	}

	sc, err := newSvcCmd(s.Name, s.Commands)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cmdHelpers := []*structTmplHelper{
		sc,
	}

	for _, c := range s.Commands {
		method, ok := field.Type.MethodByName(c.Name)
		if !ok {
			return nil, errors.New("can't find method")
		}
		cmdHelper, err := buildCommandStruct(s.Name, method.Name, method, c)
		if err != nil {
			return nil, errors.Wrap(err, "creating cmdHelper")
		}
		cmdHelpers = append(cmdHelpers, cmdHelper)
	}

	return &pkg{
		PackageName: strings.ToLower(s.Name + "svc"),
		Imports:     pkgImports,
		CmdHelpers:  cmdHelpers,
	}, nil
}

func (p *pkg) write(wr io.Writer) error {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, "svcpackage", p)
	if err != nil {
		return errors.Wrap(err, "")
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "")
	}
	out, err = imports.Process("", out, nil)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = wr.Write(out)
	return errors.Wrap(err, "")
}

func (p *pkg) writeGoFile(wr io.Writer, template string) error {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, template, p)
	if err != nil {
		return errors.Wrap(err, "")
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "")
	}
	out, err = imports.Process("", out, nil)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = wr.Write(out)
	return errors.Wrap(err, "")
}

func (s *svc) writePkg(wr io.Writer) error {
	pkg, err := s.pkg()
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = pkg.write(wr)
	return errors.Wrap(err, "")
}

func (p *pkg) outputDir(servicesBase string) string {
	return filepath.Join(servicesBase, p.PackageName)
}

func (p *pkg) writeToDir(path string) error  {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "")
	}
	f := []struct{
		filename, templateName string
	}{
		{filename: p.PackageName + ".go", templateName: "svcpackage"},
		{filename: "testhelper_test.go", templateName: "testhelper"},
	}
	for _, v := range f {
		var buf bytes.Buffer
		err = p.writeGoFile(&buf, v.templateName)
		if err != nil {
			return errors.Wrap(err, "")
		}
		fl := filepath.Join(path, v.filename)
		err = ioutil.WriteFile(fl, buf.Bytes(), 0644)
		if err != nil {
			return errors.Wrap(err, "")
		}
	}
	return nil
}

func flagName(fieldName string) string {
	s := strings.ToLower(strings.Join(camelcase.Split(fieldName), "-"))
	return strings.Replace(s, "_", "-", -1)
}

func unexport(name ...string) string {
	var words []string
	for _, v := range name {
		w := camelcase.Split(v)
		for i, ww := range w {
			w[i] = strings.Title(ww)
		}
		words = append(words, w...)
	}

	if len(words) < 1 {
		return ""
	}
	words[0] = strings.ToLower(words[0])
	return strings.Join(words, "")
}

// takes input like "foo-bar" and returns "FooBar"
func toArgName(in string) string {
	out := in
	for _, separator := range []string{"_", "-"} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}

func funcInTypes(funcType reflect.Type, offset int) []reflect.Type {
	var types []reflect.Type
	for i := offset; i < funcType.NumIn(); i++ {
		types = append(types, funcType.In(i))
	}
	return types
}

func commonStructFields() ([]structField, error) {
	var structFields []structField

	foo := []struct{
		Name string
		Type string
		Tags []*structtag.Tag
	}{
		{
			Name: "Token",
			Type: "string",
			Tags: []*structtag.Tag{
				newTag("env", "GITHUB_TOKEN"),
				newTag("required", ""),
			},
		},
		{
			Name: "APIBaseURL",
			Type: "string",
			Tags: []*structtag.Tag{
				newTag("env", "GITHUB_API_BASE_URL"),
				newTag("default", "https://api.github.com"),
			},
		},
	}

	for _, sf := range foo {
		tgs, err := newTags(sf.Tags...)
		if  err != nil {
			return nil, errors.Wrap(err, "")
		}
		structFields = append(structFields, structField{
			Name: sf.Name,
			Type: sf.Type,
			Tags: tgs,
		})
	}

	return structFields, nil
}

func buildCommandStruct(svcName, funcName string, apiFunc reflect.Method, c cmd) (*structTmplHelper, error) {
	structName := svcName + funcName + "Cmd"
	argNames := c.ArgNames
	for i, argName := range argNames {
		argNames[i] = toArgName(argName)
	}
	fullArgNames := argNames

	structFields, err := commonStructFields()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var oss []optionsStruct

	inTypes := funcInTypes(apiFunc.Type, 2)
	for _, inType := range inTypes {
		switch inType.Kind() {
		case reflect.Ptr:
			if inType.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("only pointers to structs are allowed")
			}
			oStruct, err := generateOptionsStruct(structName, inType.Elem(), c.Route)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			oss = append(oss, *oStruct)
			field := structField{
				Type: oStruct.StructName,
			}
			structFields = append(structFields, field)
		case reflect.String, reflect.Int, reflect.Bool:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			param := c.Route.Params.findByArgName(argName)
			field := newStructField(argName, inType.Kind().String(), &structtag.Tags{})
			if param != nil && param.Required {
				err := field.Tags.Set(newTag("required", ""))
				if err != nil {
					return nil, errors.Wrap(err, "")
				}
			}
			structFields = append(structFields, *field)
		case reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			param := c.Route.Params.findByArgName(argName)
			field := newStructField(argName, inType.String(), &structtag.Tags{})
			if param.Required {
				err := field.Tags.Set(newTag("required", ""))
				if err != nil {
					return nil, errors.Wrap(err, "")
				}
			}
			structFields = append(structFields, *field)
		default:
			return nil, fmt.Errorf(`buildCommandStruct: unsupported type: "%s"`, inType.Kind().String())
		}
	}
	runMethod, err := generateRunMethod(svcName, funcName, apiFunc, fullArgNames...)
	if err != nil {
		return nil, err
	}
	return &structTmplHelper{
		Name:           structName,
		Fields:         structFields,
		RunMethod:      runMethod,
		OptionsStructs: oss,
	}, nil
}

func generateRunMethod(svcName, funcName string, apiFunc reflect.Method, argNames ...string) (*runMethodHelper, error) {
	apiFuncType := apiFunc.Type
	numOut := apiFuncType.NumOut()
	structName := svcName + funcName + "Cmd"

	runStruct := &runMethodHelper{
		StructName: structName,
		FuncName:   funcName,
		SvcName:    svcName,
	}

	switch numOut {
	case 3:
		runStruct.HasElement = true
	case 2:
	default:
		return nil, fmt.Errorf("we only take funcs that return 2 or 3 args, not %d", numOut)
	}

	inTypes := funcInTypes(apiFuncType, 2)
	for _, inType := range inTypes {
		switch inType.Kind() {
		case reflect.Ptr:
			if inType.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("only pointers to structs are allowed")
			}
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{Name: inType.Elem().Name(), IsPtr: true})
		case reflect.String, reflect.Int, reflect.Bool:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{Name: argName})
		case reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{
				Name: argName,
			})
		default:
			return nil, fmt.Errorf("generateRunMethod: unsupported type: %v", inType.Kind())
		}
	}

	return runStruct, nil
}

func optionsStructName(methodName string, requestType reflect.Type) string {
	return unexport(methodName, requestType.Name()+"Flags")
}

func typeToFields(inType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < inType.NumField(); i++ {
		field := inType.Field(i)
		if field.Anonymous {
			ft := field.Type
			if ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			fields = append(fields, typeToFields(ft)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func generateOptionsStruct(cmdName string, requestType reflect.Type, route *svcRoute) (*optionsStruct, error) {
	structName := optionsStructName(cmdName, requestType)

	fields := typeToFields(requestType)
	structFields, err := getStructFields(fields, route)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var keeperFields []reflect.StructField
	for _, field := range fields {
		for _, sf := range structFields {
			if sf.Name == field.Name {
				keeperFields = append(keeperFields, field)
			}
		}
	}

	oStruct := optionsStruct{
		StructName: structName,
		MainStruct: structTmplHelper{
			Name:   structName,
			Fields: structFields,
		},
		ToFunc: generateToRequestFunc(keeperFields, structName, requestType),
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "options_struct", &oStruct)
	return &oStruct, err
}

func getStructFields(fields []reflect.StructField, route *svcRoute) ([]structField, error) {
	var structFields []structField
	for _, field := range fields {
		if field.Type.Kind() == reflect.Ptr {
			field.Type = field.Type.Elem()
		}
		tags, err := newTags(newTag("name", flagName(field.Name)))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		if route != nil {
			param := route.Params.forField(field)
			if param == nil {
				continue
			}
			if param.Location != "body" {
				continue
			}
			if param != nil {
				if param.Required {
					err := tags.Set(newTag("required", ""))
					if err != nil {
						return nil, errors.Wrap(err, "")
					}
				}

				if param.Description != "" {
					err := tags.Set(newTag("help", param.Description))
					if err != nil {
						return nil, errors.Wrap(err, "")
					}
				}
			}
		}
		sField := newStructField(field.Name, field.Type.String(), tags)
		structFields = append(structFields, *sField)

	}
	return structFields, nil
}

func (p *routeParams) forField(f reflect.StructField) *routeParam {
	jsonTag := f.Tag.Get("json")
	var name string

	if jsonTag == "" {
		name = strings.Join(camelcase.Split(f.Name), "-")
	} else {
		name = strings.Split(jsonTag, ",")[0]
	}

	return p.findByName(name)
}

func fieldFlagName(f reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	var name string

	if jsonTag == "" {
		name = strings.Join(camelcase.Split(f.Name), "-")
	} else {
		name = strings.Split(jsonTag, ",")[0]
	}
	return strings.ToLower(strings.Replace(name, "_", "-", -1))
}

func generateToRequestFunc(fields []reflect.StructField, structName string, targetType reflect.Type) toFuncTmplHelper {
	inclPtrHelper := false
	for _, v := range fields {
		if v.Type.Kind() == reflect.Ptr {
			inclPtrHelper = true
			break
		}
	}

	var vss []valSetter
	for _, field := range fields {
		vss = append(vss, valSetter{
			Name:        field.Name,
			FlagName:    fieldFlagName(field),
			TargetIsPtr: field.Type.Kind() == reflect.Ptr,
		})
	}

	return toFuncTmplHelper{
		ReceiverName:         structName,
		TargetName:           targetType.Name(),
		TargetType:           targetType.String(),
		IncludePointerHelper: inclPtrHelper,
		ValSetters:           vss,
	}
}

func (p *routeParams) findByArgName(name string) *routeParam {
	for _, param := range *p {
		if toArgName(param.Name) == toArgName(name) {
			return param
		}
	}
	return nil
}

func (p *routeParams) findByName(name string) *routeParam {
	for _, param := range *p {
		if param.Name == name {
			return param
		}
	}
	return nil
}

func (r *svcRoutes) findByIdName(idName string) *svcRoute {
	for _, route := range *r {
		if route.IDName == idName {
			return route
		}
	}
	return nil
}

func getRtServices(file string) (*rtServices, error) {
	var sm rtServices
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	err = json.Unmarshal(bts, &sm)
	return &sm, errors.Wrap(err, "")
}

func (r rtServices) svcRoutes(serviceName string) *svcRoutes {
	sr := r[serviceName]
	return &sr
}

// language=GoTemplate
const pkgTemplate = `
{{define "run_method_arg"}}{{if .IsPtr}}c.to{{.Name}}(k){{else}}c.{{.Name}}{{end}}{{end}}

{{define "run_method_args"}}{{range $index, $element := .Args}}{{if $index}}, {{end}}{{template "run_method_arg" .}}{{end}}{{end}}

{{define "run_method"}}
	{{if .}}
	func (c *{{.StructName}}) Run(k *kong.Context) error {
			ctx := context.Background()
			client, e := buildGithubClient(ctx, c.Token, c.APIBaseURL)
			if e != nil {
				return e
			}
			{{if .HasElement}}element, {{end}}_, err := client.{{.SvcName}}.{{.FuncName}}(ctx, {{template "run_method_args" .}})
			{{if .HasElement}}	if err != nil {
			return err
		}
		return json.NewEncoder(k.Stdout).Encode(element){{else}}	return err{{end}}
	}
	{{end}}
{{end}}

{{define "structtype"}}
	type {{.Name}} struct { {{range .Fields}}
		{{.Name}} {{.Type}} {{if .Tags}}{{printf "%#q" .Tags}} {{end}}{{end}}
	}
	{{template "run_method" .RunMethod}}{{template "options_structs" .OptionsStructs}}
{{end}}

{{define "tofunc"}}
	func (t {{.ReceiverName}}) to{{.TargetName}}(k *kong.Context) *{{.TargetType}} {
		val := &{{.TargetType}}{}
		{{if .IncludePointerHelper}}
			isValueSet := func (valueName string) bool {
				if k == nil {
					return false
				}
				for _, flag := range k.Flags() {
					if flag.Name == valueName {
						return flag.Set
					}
				}
				return false
			}
		{{end}}
		{{template "val_setters" .ValSetters}}
		return val
	}
{{end}}

{{define "val_setters"}}{{range .}}{{template "val_setter" .}}{{end}}{{end}}

{{define "val_setter"}}{{if .TargetIsPtr}}	if isValueSet("{{.FlagName}}") {
		val.{{.Name}} = &t.{{.Name}}
	}{{else}}	val.{{.Name}} = t.{{.Name}}{{end}}

{{end}}

{{define "options_struct"}}
{{template "structtype" .MainStruct}}

{{template "tofunc" .ToFunc}}
{{end}}

{{define "options_structs"}}
{{range .}}{{template "options_struct" .}}{{end}}
{{end}}

{{define "svcpackage"}}
	// Code generated by go-github-cli/generator DO NOT EDIT
	package {{$.PackageName}}

	import ( {{range .Imports}}
	   "{{.}}"{{end}}
	)

	var transportWrapper interface {
		SetTransport(t http.RoundTripper)
		http.RoundTripper
	}
	
	func buildGithubClient(ctx context.Context, token, apiBaseURL string) (*github.Client, error) {
		apiBaseURL  = strings.TrimSuffix(apiBaseURL, "/") + "/"
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		if transportWrapper != nil {
			transportWrapper.SetTransport(tc.Transport)
			tc.Transport = transportWrapper
		}
		client := github.NewClient(tc)
		baseURL, err := url.Parse(apiBaseURL)
		client.BaseURL = baseURL
		return client, err
	}

	{{range .CmdHelpers}}{{template "structtype" .}}{{end}}
	{{range .OptionStructs}}{{template "options_struct" .}}{{end}}
{{end}}

{{define "testhelper"}}
	// Code generated by go-github-cli/generator DO NOT EDIT
	package {{$.PackageName}}

	import (
		"bytes"
		"github.com/alecthomas/kong"
		"github.com/dnaeon/go-vcr/recorder"
		"github.com/stretchr/testify/require"
		"os"
		"path/filepath"
		"testing"
	)
	
	func init() {
		tkn, ok := os.LookupEnv("TESTUSER_TOKEN")
		if !ok {
			tkn = "deadbeef"
		}
		os.Setenv("GITHUB_TOKEN", tkn)
	}
	
	func startVCR(t *testing.T, recPath string) *recorder.Recorder {
		t.Helper()
		var err error
		rec, err := recorder.New(recPath)
		require.Nil(t, err)
		transportWrapper = rec
		return rec
	}
	
	func testCmdLine(t *testing.T, fixtureName string, cmdStruct interface{}, cmdline ...string) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
		t.Helper()
		rec := startVCR(t, filepath.Join("testdata", "fixtures", fixtureName))
		defer rec.Stop()
		p, e := kong.New(cmdStruct)
		require.Nil(t, e)
		p.Stdout = &stdout
		p.Stderr = &stderr
		k, e := p.Parse(cmdline)
		require.Nil(t, e)
		err = k.Run()
		return
	}
{{end}}

`
