package overrides

import (
	"github.com/dave/jennifer/jen"
	"github.com/octo-cli/octo-cli/internal/model"
)

func OverrideEndpoints(endpoints []model.Endpoint) {
	for i := 0; i < len(endpoints); i++ {
		endpoint := &endpoints[i]
		if endpointOverrides[endpoint.ID] == nil {
			continue
		}
		endpointOverrides[endpoint.ID](endpoint)
		endpoints[i] = *endpoint
	}
}

// no overrides are currently necessary, but if there were, they would go here.
var endpointOverrides = map[string]func(endpoint *model.Endpoint){}

type ManualParamInfo struct {
	Name        string
	Type        string
	CodeAdder   func(group *jen.Group)
	CodeImports []string
	FieldImport string
	Required    bool
	Description string
	Tags        map[string]string
	Hidden      bool
}

func GetManualParamInfo(opID string) []ManualParamInfo {
	overrides := map[string][]ManualParamInfo{
		"markdown/render-raw": {
			{
				Name:        "file",
				Type:        "string",
				Tags:        map[string]string{"type": "existingfile"},
				Required:    true,
				CodeImports: []string{"github.com/octo-cli/octo-cli/internal"},
				Description: "the file to upload",
				CodeAdder: func(group *jen.Group) {
					group.Qual("github.com/octo-cli/octo-cli/internal", "MarkdownRenderRawOverride").
						Params(jen.Id("&c.BaseCmd"), jen.Id("c.File"))
				},
			},
			{
				Name:   "content-type",
				Type:   "string",
				Tags:   map[string]string{"hidden": ""},
				Hidden: true,
			},
		},
		"repos/upload-release-asset": {
			{
				Name:        "file",
				Type:        "string",
				Tags:        map[string]string{"type": "existingfile"},
				Required:    true,
				CodeImports: []string{"github.com/octo-cli/octo-cli/internal"},
				Description: "the file to upload",
				CodeAdder: func(group *jen.Group) {
					group.Qual("github.com/octo-cli/octo-cli/internal", "ReposUploadReleaseAssetOverride").
						Params(jen.Id("&c.BaseCmd"), jen.Id("c.File"))
				},
			},
			{
				Name:        "content-type",
				Type:        "string",
				Description: "override the Content-Type header",
			},
			{
				Name:   "content-length",
				Type:   "string",
				Tags:   map[string]string{"hidden": ""},
				Hidden: true,
			},
		},
	}
	return overrides[opID]
}
