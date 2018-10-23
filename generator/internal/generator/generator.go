package generator

import (
	"fmt"
	"github.com/WillAbides/go-github-cli/generator/internal"
	"github.com/WillAbides/go-github-cli/generator/internal/configparser"
	"github.com/WillAbides/go-github-cli/generator/internal/routeparser"
	"github.com/fatih/camelcase"
	"github.com/fatih/structtag"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"reflect"
	"sort"
	"strings"
)

type (
	// Svc represents a group of api endpoints such as Issues, Organizations or Git
	Svc struct {
		Name     string
		Commands []*cmd
	}

	// cmd represents a cli command that will be generated
	cmd struct {
		Name     string
		ArgNames []string
		Route    *routeparser.Route
	}
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
)

//BuildSvcs builds services
func BuildSvcs(routesPath, configFile string) ([]Svc, error) {
	rt, err := routeparser.ParseRoutesFile(routesPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parsing routes at %q", routesPath)
	}
	hConfig, err := configparser.ParseConfigFile(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parsing config file at %q", configFile)
	}

	var svcs []Svc

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
		cmds := make([]*cmd, len(cmdNames))
		for i, cmdName := range cmdNames {
			hcmd := hsvc.Command[cmdName]
			routesName := hcmd.RoutesName
			if routesName == "" {
				routesName = flagName(cmdName)
			}
			routes := rt[svcRoutesName]
			route := routes.FindByIDName(routesName)
			cmds[i] = newCmd(cmdName, route, hcmd.ArgNames...)
		}
		svcs = append(svcs, Svc{
			Name:     svcName,
			Commands: cmds,
		})
	}
	return svcs, nil
}

//newCmd creates a new cmd.  If argNames is empty, it infers them from them from url parameters
func newCmd(name string, rt *routeparser.Route, argNames ...string) *cmd {
	if len(argNames) == 0 {
		for _, p := range rt.Params {
			if p.Location == "url" {
				argNames = append(argNames, internal.ToArgName(p.Name))
			}
		}
	}

	return &cmd{
		Name:     name,
		ArgNames: argNames,
		Route:    rt,
	}
}

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

//appendTag appends one of more tags to a *structtag.Tags
func appendTag(tags *structtag.Tags, tag ...*structtag.Tag) *structtag.Tags {
	return newTags(append(tags.Tags(), tag...)...)
}

//newSvcCmd create a structTmplHelper that represents the top level command of a service
func newSvcCmd(svcName string, cmds []*cmd) *internal.StructTmplHelper {
	var fields []internal.StructField
	for _, cmd := range cmds {
		fields = append(fields, internal.StructField{
			Name: cmd.Name,
			Type: svcName + cmd.Name + "Cmd",
			Tags: newTags(newTag("cmd", ""), newTag("help", cmd.Route.Name)),
		})
	}
	return &internal.StructTmplHelper{
		Name:   svcName + "Cmd",
		Fields: fields,
	}
}

func (s *Svc) getStructField() (reflect.StructField, bool) {
	clientType := reflect.TypeOf(github.Client{})
	return clientType.FieldByName(s.Name)
}

func (s Svc) ToPkg() (*internal.Pkg, error) {
	field, ok := s.getStructField()
	if !ok {
		return nil, errors.New("can't find structField")
	}

	cmdHelpers := []*internal.StructTmplHelper{
		newSvcCmd(s.Name, s.Commands),
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

	return &internal.Pkg{
		PackageName: strings.ToLower(s.Name + "svc"),
		Imports:     pkgImports,
		CmdHelpers:  cmdHelpers,
	}, nil
}

func flagName(fieldName string) string {
	s := strings.ToLower(strings.Join(camelcase.Split(fieldName), "-"))
	return strings.Replace(s, "_", "-", -1)
}

func funcInTypes(funcType reflect.Type, offset int) []reflect.Type {
	var types []reflect.Type
	for i := offset; i < funcType.NumIn(); i++ {
		types = append(types, funcType.In(i))
	}
	return types
}

func buildCommandStruct(svcName, funcName string, apiFunc reflect.Method, c *cmd) (*internal.StructTmplHelper, error) {
	structName := svcName + funcName + "Cmd"
	argNames := c.ArgNames
	for i, argName := range argNames {
		argNames[i] = internal.ToArgName(argName)
	}
	fullArgNames := argNames

	structFields := []internal.StructField{
		{
			Name: "Token",
			Type: "string",
			Tags: newTags(newTag("env", "GITHUB_TOKEN"), newTag("required", "")),
		},
		{
			Name: "APIBaseURL",
			Type: "string",
			Tags: newTags(newTag("env", "GITHUB_API_BASE_URL"), newTag("default", "https://api.github.com")),
		},
	}

	var oss []internal.StructTmplHelper

	inTypes := funcInTypes(apiFunc.Type, 2)
	for _, inType := range inTypes {
		switch inType.Kind() {
		case reflect.Ptr:
			if inType.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("only pointers to structs are allowed")
			}
			oStruct := generateOptionsStruct(structName, inType.Elem(), c.Route)
			oss = append(oss, *oStruct)
			structFields = append(structFields, internal.StructField{
				Type: optionsStructName(structName, inType.Elem()),
			})
		case reflect.String, reflect.Int, reflect.Bool, reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			param := c.Route.ArgParam(argName)
			field := internal.StructField{
				Name: argName,
				Type: inType.String(),
				Tags: &structtag.Tags{},
			}
			if param != nil && param.Required {
				field.Tags = appendTag(field.Tags, newTag("required", ""))
			}
			structFields = append(structFields, field)
		default:
			return nil, fmt.Errorf(`buildCommandStruct: unsupported type: "%s"`, inType.Kind().String())
		}
	}
	runMethod, err := generateRunMethod(svcName, funcName, apiFunc, fullArgNames...)
	if err != nil {
		return nil, err
	}
	return &internal.StructTmplHelper{
		Name:         structName,
		Fields:       structFields,
		RunMethod:    runMethod,
		ChildStructs: oss,
	}, nil
}

func generateRunMethod(svcName, funcName string, apiFunc reflect.Method, argNames ...string) (*internal.RunMethod, error) {
	apiFuncType := apiFunc.Type
	numOut := apiFuncType.NumOut()
	structName := svcName + funcName + "Cmd"

	runStruct := &internal.RunMethod{
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
			runStruct.Args = append(runStruct.Args, internal.RunMethodArg{Name: inType.Elem().Name(), IsPtr: true})
		case reflect.String, reflect.Int, reflect.Bool, reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			runStruct.Args = append(runStruct.Args, internal.RunMethodArg{Name: argName})
		default:
			return nil, fmt.Errorf("generateRunMethod: unsupported type: %v", inType.Kind())
		}
	}

	return runStruct, nil
}

func optionsStructName(methodName string, requestType reflect.Type) string {
	return internal.Unexport(methodName, requestType.Name()+"Flags")
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

func generateOptionsStruct(cmdName string, requestType reflect.Type, route *routeparser.Route) *internal.StructTmplHelper {
	structName := optionsStructName(cmdName, requestType)

	fields := typeToFields(requestType)
	structFields := getStructFields(fields, route)

	var keeperFields []reflect.StructField
	for _, field := range fields {
		for _, sf := range structFields {
			if sf.Name == field.Name {
				keeperFields = append(keeperFields, field)
			}
		}
	}
	return &internal.StructTmplHelper{
		Name:   structName,
		Fields: structFields,
		ToFunc: generateToRequestFunc(keeperFields, structName, requestType),
	}
}

func getStructFields(fields []reflect.StructField, route *routeparser.Route) []internal.StructField {
	var structFields []internal.StructField
	for _, field := range fields {
		if field.Type.Kind() == reflect.Ptr {
			field.Type = field.Type.Elem()
		}
		tags := newTags(newTag("name", flagName(field.Name)))
		if route != nil {
			param := route.FieldParam(field)
			if param == nil {
				continue
			}
			if param.Location != "body" {
				continue
			}
			if param.Required {
				tags = appendTag(tags, newTag("required", ""))
			}

			if param.Description != "" {
				tags = appendTag(tags, newTag("help", param.Description))
			}
		}
		structFields = append(structFields, internal.StructField{
			Name: field.Name,
			Type: field.Type.String(),
			Tags: tags,
		})

	}
	return structFields
}

//fieldFlagName gets the flag name for a field either from its json tag or its name
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

func generateToRequestFunc(fields []reflect.StructField, structName string, targetType reflect.Type) *internal.ToFunc {
	var valSetters []internal.ValSetter
	for _, field := range fields {
		valSetters = append(valSetters, internal.ValSetter{
			Name:        field.Name,
			FlagName:    fieldFlagName(field),
			TargetIsPtr: field.Type.Kind() == reflect.Ptr,
		})
	}

	return &internal.ToFunc{
		ReceiverName: structName,
		TargetName:   targetType.Name(),
		TargetType:   targetType.String(),
		ValSetters:   valSetters,
	}
}
