package generator

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type (
	//Route represents one route from routes.json.  Such as "issues edit"
	Route struct {
		Description      string        `json:"description,omitempty"`
		Method           string        `json:"method,omitempty"`
		Path             string        `json:"path,omitempty"`
		Name             string        `json:"name,omitempty"`
		EnabledForApps   bool          `json:"enabledForApps,omitempty"`
		IDName           string        `json:"idName,omitempty"`
		DocumentationURL string        `json:"documentationUrl,omitempty"`
		Params           []*Param      `json:"params,omitempty"`
		Requests         []interface{} `json:"requests,omitempty"`
		Previews         []*Preview    `json:"previews,omitempty"`
	}

	//Param represents one parameter for a Route such as repo name or issue number
	Param struct {
		Name        string        `json:"name,omitempty"`
		Type        string        `json:"type,omitempty"`
		Description string        `json:"description,omitempty"`
		Default     interface{}   `json:"default,omitempty"`
		Required    bool          `json:"required,omitempty"`
		Enum        []interface{} `json:"enum,omitempty"`
		Location    string        `json:"location,omitempty"`
		MapTo       string        `json:"mapTo,omitempty"`
	}

	//Routes is a collection or Routes
	Routes []*Route

	//Preview is a preview header
	Preview struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Required    bool   `json:"required,omitempty"`
	}
)

type xGitHub struct {
	Legacy          bool
	EnabledForApps  bool
	GithubCloudOnly bool
	Previews        []xGitHubPreview
}

type xGitHubPreview struct {
	Name     string
	Required bool
	Note     string
}

func parseRoutesFile(file string) (map[string]Routes, error) {
	sm := map[string]Routes{}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(file)
	if err != nil {
		return nil, err
	}
	for path, item := range swagger.Paths {
		path = strings.ReplaceAll(path, "{", ":")
		path = strings.ReplaceAll(path, "}", "")
		for method, op := range item.Operations() {
			opID := strings.Split(op.OperationID, "/")
			svc := dashCamel(opID[0])
			idName := "invalid"
			if len(opID) > 1 {
				idName = opID[1]
			}
			xg := xGitHub{}
			xgMsg, ok := op.Extensions["x-github"].(json.RawMessage)
			if ok {
				err = json.Unmarshal(xgMsg, &xg)
				if err != nil {
					return nil, err
				}
			}
			var previews []*Preview
			for _, preview := range xg.Previews {
				previews = append(previews, &Preview{
					Name:        preview.Name,
					Description: preview.Note,
					Required:    preview.Required,
				})
			}
			params := append(opParams(op), bodyParams(op)...)

			route := Route{
				Path:             path,
				Method:           method,
				Description:      op.Description,
				IDName:           idName,
				Name:             op.Summary,
				DocumentationURL: op.ExternalDocs.URL,
				Previews:         previews,
				EnabledForApps:   xg.EnabledForApps,
				Params:           params,
			}
			sm[svc] = append(sm[svc], &route)
		}
	}
	return sm, nil
}

func opParams(op *openapi3.Operation) []*Param {
	var params []*Param
	for _, pRef := range op.Parameters {
		param := pRef.Value
		if param.Name == "accept" {
			continue
		}
		location := param.In
		if location == "path" {
			location = "url"
		}
		p := &Param{
			Name:        param.Name,
			Type:        param.Schema.Value.Type,
			Description: param.Description,
			Default:     param.Schema.Value.Default,
			Required:    param.Required,
			Enum:        param.Schema.Value.Enum,
			Location:    location,
		}
		params = append(params, p)
	}
	return params
}

func bodyParams(op *openapi3.Operation) []*Param {
	var params []*Param
	if op.RequestBody == nil {
		return params
	}
	for s, mediaType := range op.RequestBody.Value.Content {
		if s != "application/json" {
			continue
		}
		required := map[string]bool{}
		for _, req := range mediaType.Schema.Value.Required {
			required[req] = true
		}
		for nm, prop := range mediaType.Schema.Value.Properties {
			p := &Param{
				Name:        nm,
				Required:    required[nm],
				Location:    "body",
				Type:        getPropType(prop.Value),
				Enum:        prop.Value.Enum,
				Default:     prop.Value.Default,
				Description: prop.Value.Description,
			}
			params = append(params, p)
		}
	}
	return params
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

func dashCamel(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	runes := []rune(generator.CamelCase(s))
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}
