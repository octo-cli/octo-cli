package openapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/octo-cli/octo-cli/internal/model"
	"github.com/pkg/errors"
)

func EndpointsFromSchema(schemaSrc io.Reader) ([]model.Endpoint, error) {
	data, err := ioutil.ReadAll(schemaSrc)
	if err != nil {
		return nil, fmt.Errorf("could not read from schemaSrc")
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	if err != nil {
		return nil, fmt.Errorf("could not load openapiDef")
	}
	return buildEndpoints(swagger)
}

func buildEndpoints(swagger *openapi3.Swagger) ([]model.Endpoint, error) {
	var endpoints []model.Endpoint
	for opPath, pathItem := range swagger.Paths {
		for method, op := range pathItem.Operations() {
			endpoint, err := buildEndpoint(opPath, method, op)
			if err != nil {
				return nil, err
			}
			endpoints = append(endpoints, *endpoint)
		}
	}
	return endpoints, nil
}

func buildEndpoint(opPath, httpMethod string, op *openapi3.Operation) (*model.Endpoint, error) {
	endpoint := model.Endpoint{
		Path:       opPath,
		Method:     httpMethod,
		Name:       path.Base(op.OperationID),
		Concern:    path.Dir(op.OperationID),
		Summary:    op.Summary,
		HelpText:   op.Description,
		Deprecated: op.Deprecated,
		ID:         op.OperationID,
	}
	ext, err := opExtGithub(op)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if ext != nil {
		endpoint.EnabledForApps = ext.EnabledForApps
		endpoint.GithubCloudOnly = ext.GithubCloudOnly
		endpoint.Legacy = ext.Legacy
		for _, preview := range ext.Previews {
			endpoint.Previews = append(endpoint.Previews, model.Preview{
				Required: preview.Required,
				Name:     preview.Name,
				Note:     preview.Note,
			})
		}
	}
	if op.ExternalDocs != nil {
		endpoint.DocsURL = op.ExternalDocs.URL
	}
	for _, pRef := range op.Parameters {
		var param *model.Param
		param, err = buildParam(pRef.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		switch pRef.Value.In {
		case openapi3.ParameterInQuery:
			endpoint.QueryParams = append(endpoint.QueryParams, *param)
		case openapi3.ParameterInHeader:
			endpoint.Headers = append(endpoint.Headers, *param)
		case openapi3.ParameterInPath:
			endpoint.PathParams = append(endpoint.PathParams, *param)
		}
	}
	endpoint.JSONBodySchema, err = jsonBodySchema(op)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error processing operation %q", op.OperationID))
	}
	return &endpoint, nil
}

var op2modelTypes = map[string]model.ParamType{
	"string":  model.ParamTypeString,
	"":        model.ParamTypeString,
	"number":  model.ParamTypeString,
	"integer": model.ParamTypeInt,
	"boolean": model.ParamTypeBool,
	"object":  model.ParamTypeObject,
	"array":   model.ParamTypeArray,
}

func opSchemaType(opSchema *openapi3.Schema) model.ParamType {
	if strings.HasPrefix(opSchema.Type, "[]") {
		return model.ParamTypeArray
	}
	return op2modelTypes[opSchema.Type]
}

func opParamSchema(opSchema *openapi3.Schema) (*model.ParamSchema, error) {
	schema := model.ParamSchema{
		Type: opSchemaType(opSchema),
	}
	var err error
	switch schema.Type {
	case model.ParamTypeInvalid:
		return nil, fmt.Errorf("unknown schema type %s", opSchema.Type)
	case model.ParamTypeArray:
		schema.ItemSchema, err = opParamSchema(opSchema.Items.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	case model.ParamTypeObject:
		if opSchema.AdditionalProperties != nil {
			schema.ItemSchema, err = opParamSchema(opSchema.AdditionalProperties.Value)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
		propNames := make([]string, 0, len(opSchema.Properties))
		for name := range opSchema.Properties {
			propNames = append(propNames, name)
		}
		sort.Strings(propNames)
		for _, name := range propNames {
			ref := opSchema.Properties[name]
			var objParam *model.Param
			objParam, err = opObjectParam(opSchema, ref.Value, name)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			schema.ObjectParams = append(schema.ObjectParams, *objParam)
		}
	}
	return &schema, nil
}

func opObjectParam(opSchema, propSchema *openapi3.Schema, name string) (*model.Param, error) {
	param := model.Param{
		Name:     name,
		HelpText: propSchema.Description,
	}
	for _, reqName := range opSchema.Required {
		if name == reqName {
			param.Required = true
			break
		}
	}
	var err error
	param.Schema, err = opParamSchema(propSchema)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &param, nil
}

func buildParam(opParam *openapi3.Parameter) (*model.Param, error) {
	schema, err := opParamSchema(opParam.Schema.Value)
	if err != nil {
		return nil, err
	}
	param := model.Param{
		Required: opParam.Required,
		Name:     opParam.Name,
		HelpText: opParam.Description,
		Schema:   schema,
	}
	return &param, nil
}

func opExtGithub(op *openapi3.Operation) (*extGithub, error) {
	xMsg, ok := op.Extensions["x-github"].(json.RawMessage)
	if !ok {
		return nil, nil
	}
	var ext extGithub
	err := json.Unmarshal(xMsg, &ext)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ext, nil
}

type extGithub struct {
	Legacy          bool
	EnabledForApps  bool
	GithubCloudOnly bool
	Previews        []struct {
		Name     string
		Required bool
		Note     string
	}
}

func jsonBodySchema(op *openapi3.Operation) (*model.ParamSchema, error) {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil, nil
	}
	mt := op.RequestBody.Value.GetMediaType("application/json")
	if mt == nil {
		return nil, nil
	}
	schema, err := opParamSchema(mt.Schema.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return schema, nil
}
