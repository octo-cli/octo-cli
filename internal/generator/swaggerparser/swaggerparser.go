package swaggerparser

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/octo-cli/octo-cli/internal/generator/util"
)

func ForEachOperation(swagger *openapi3.Swagger, fn func(path, method string, op *openapi3.Operation)) {
	//nolint:errcheck
	_ = ForEachOperationErr(swagger, func(path, method string, op *openapi3.Operation) error {
		fn(path, method, op)
		return nil
	})
}

func ForEachOperationErr(swagger *openapi3.Swagger, fn func(path, method string, op *openapi3.Operation) error) error {
	for path, pathItem := range swagger.Paths {
		for method, op := range pathItem.Operations() {
			err := fn(path, method, op)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetOperationSvcName(op *openapi3.Operation) string {
	return strings.Split(op.OperationID, "/")[0]
}

func GetOperationName(op *openapi3.Operation) string {
	return strings.TrimPrefix(op.OperationID, GetOperationSvcName(op)+"/")
}

type Preview struct {
	Name     string
	Required bool
	Note     string
}

func OperationPreviews(op *openapi3.Operation) ([]Preview, error) {
	xMsg, ok := op.Extensions["x-github"].(json.RawMessage)
	if !ok {
		return nil, nil
	}
	xg := struct {
		Legacy          bool
		EnabledForApps  bool
		GithubCloudOnly bool
		Previews        []Preview
	}{}
	err := json.Unmarshal(xMsg, &xg)
	if err != nil {
		return nil, err
	}
	for i := range xg.Previews {
		xg.Previews[i].Note = strings.TrimSpace(xg.Previews[i].Note)
		xg.Previews[i].Note = strings.Split(xg.Previews[i].Note, "```")[0]
		xg.Previews[i].Note = strings.TrimSpace(xg.Previews[i].Note)
		setThisFlagPhrases := []string{
			"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
			"provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` Header",
			"provide the following custom [media type](https://developer.github.com/v3/media) in the `Accept` header",
		}
		for _, phrase := range setThisFlagPhrases {
			xg.Previews[i].Note = strings.ReplaceAll(xg.Previews[i].Note, phrase, "set this flag")
		}
		xg.Previews[i].Note = strings.TrimSpace(xg.Previews[i].Note)
		xg.Previews[i].Note = strings.TrimSuffix(xg.Previews[i].Note, ":")
		xg.Previews[i].Note = strings.TrimSpace(xg.Previews[i].Note)
		xg.Previews[i].Note = strings.TrimSuffix(xg.Previews[i].Note, ".") + "."
	}
	return xg.Previews, nil
}

func AllServiceNames(swagger *openapi3.Swagger) []string {
	mp := map[string]struct{}{}
	ForEachOperation(swagger, func(path, method string, op *openapi3.Operation) {
		mp[GetOperationSvcName(op)] = struct{}{}
	})
	result := make([]string, 0, len(mp))
	for k := range mp {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}

type BodyParamInfo struct {
	Name     string
	Ref      *openapi3.SchemaRef
	Required bool
}

func GetObjParamInfo(ref *openapi3.SchemaRef, names []string, parentRequired bool, refFilter func(schemaRef *openapi3.SchemaRef) bool) []BodyParamInfo {
	var result []BodyParamInfo
	if len(ref.Value.Properties) == 0 {
		result = append(result, BodyParamInfo{
			Name:     strings.Join(names, "."),
			Ref:      ref,
			Required: parentRequired,
		})
	}
	for nm, sr := range ref.Value.Properties {
		if refFilter != nil && !refFilter(sr) {
			continue
		}
		var required bool
		if parentRequired {
			for _, s := range ref.Value.Required {
				if s == nm {
					required = true
					break
				}
			}
		}
		if sr.Value.Type != "object" {
			result = append(result, BodyParamInfo{
				Name:     strings.Join(append(names, nm), "."),
				Ref:      sr,
				Required: required,
			})
			continue
		}
		result = append(result, GetObjParamInfo(sr, append(names, nm), required, refFilter)...)
	}
	return result
}

type ManualParamInfo struct {
	Name        string
	Type        string
	RunCode     string
	CodeImports []string
	FieldImport string
	Required    bool
	Description string
	Tags        *structtag.Tags
	Hidden      bool
}

func GetManualParamInfo(op *openapi3.Operation) []ManualParamInfo {
	overrides := map[string][]ManualParamInfo{
		"markdown/render-raw": {
			{
				Name:        "file",
				Type:        "string",
				Tags:        util.NewTags(util.NewTag("type", "existingfile")),
				Required:    true,
				CodeImports: []string{"github.com/octo-cli/octo-cli/internal"},
				Description: "the file to upload",
				RunCode: `
internal.MarkdownRenderRawOverride(&c.BaseCmd, c.File)`,
			},
			{
				Name:   "content-type",
				Type:   "string",
				Tags:   util.NewTags(util.NewTag("hidden", "")),
				Hidden: true,
			},
		},
		"repos/upload-release-asset": {
			{
				Name:        "file",
				Type:        "string",
				Tags:        util.NewTags(util.NewTag("type", "existingfile")),
				Required:    true,
				CodeImports: []string{"github.com/octo-cli/octo-cli/internal"},
				Description: "the file to upload",
				RunCode: `
internal.ReposUploadReleaseAssetOverride(&c.BaseCmd, c.File)`,
			},
			{
				Name:        "content-type",
				Type:        "string",
				Description: "override the Content-Type header",
			},
			{
				Name:   "content-length",
				Type:   "string",
				Tags:   util.NewTags(util.NewTag("hidden", "")),
				Hidden: true,
			},
		},
	}
	return overrides[op.OperationID]
}

func GetBodyParamInfo(op *openapi3.Operation, filter func(ref *openapi3.SchemaRef) bool) []BodyParamInfo {
	if op.RequestBody == nil {
		return nil
	}
	content := op.RequestBody.Value.Content.Get("application/json")
	if content == nil {
		return nil
	}
	var result []BodyParamInfo

	schemaVal := content.Schema.Value
	for name, ref := range schemaVal.Properties {
		var required bool
		for _, s := range schemaVal.Required {
			if s == name {
				required = true
				break
			}
		}
		if ref.Value.Type != "object" {
			result = append(result, BodyParamInfo{
				Name:     name,
				Ref:      ref,
				Required: required,
			})
			continue
		}
		result = append(result, GetObjParamInfo(ref, []string{name}, required, filter)...)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

func ParamRequired(parameters openapi3.Parameters, idx int) bool {
	param := parameters[idx].Value
	if param.Name != "owner" || param.In != "path" || idx == len(parameters)-1 {
		return param.Required
	}
	nextParam := parameters[idx+1].Value
	return nextParam.Name != "repo" || nextParam.In != "path"
}
