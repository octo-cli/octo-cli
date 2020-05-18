package overrides

import (
	"github.com/fatih/structtag"
	"github.com/octo-cli/octo-cli/internal/generator/util"
)

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

func GetManualParamInfo(opID string) []ManualParamInfo {
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
	return overrides[opID]
}
