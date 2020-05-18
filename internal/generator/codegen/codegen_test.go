package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/stretchr/testify/require"
)

func execAndFormatTmpl(t *testing.T, tm *template.Template, data interface{}) string {
	t.Helper()
	var buf bytes.Buffer
	err := tm.Execute(&buf, data)
	require.NoError(t, err)
	got, err := format.Source(buf.Bytes())
	require.NoError(t, err)
	return string(got)
}

func TestTemplate(t *testing.T) {
	t.Run("RunMethod", func(t *testing.T) {

		t.Run("with params", func(t *testing.T) {
			data := RunMethod{
				ReceiverName: "receiverName",
				Method:       "mymethod",
				URLPath:      "urlPath",
				CodeBlocks: []CodeBlock{
					{
						Code: `
	c.updateMethod1("param1", c.valueField1)`,
					},
					{
						Code: `
	c.updateMethod2("param2", c.valueField2)`,
					},
				},
			}
			want := `
func (c *receiverName) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("urlPath")
	c.updateMethod1("param1", c.valueField1)
	c.updateMethod2("param2", c.valueField2)
	return c.DoRequest("mymethod")
}

`
			got := execAndFormatTmpl(t, tmpl.Lookup("RunMethod"), data)
			require.Equal(t, want, got)
		})

		t.Run("no params", func(t *testing.T) {
			data := RunMethod{
				ReceiverName: "receiverName",
				Method:       "mymethod",
				URLPath:      "urlPath",
			}
			want := `
func (c *receiverName) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("urlPath")
	return c.DoRequest("mymethod")
}

`
			got := execAndFormatTmpl(t, tmpl.Lookup("RunMethod"), data)
			require.Equal(t, want, got)
		})
	})

	t.Run("StructType", func(t *testing.T) {
		t.Run("no fields", func(t *testing.T) {
			data := StructTmplHelper{
				Name: "myname",
			}
			want := `type myname struct {
}`
			got := execAndFormatTmpl(t, tmpl.Lookup("StructType"), data)
			require.Equal(t, want, got)
		})

		t.Run("with fields", func(t *testing.T) {
			fmt.Println(t.Name())
			data := StructTmplHelper{
				Name: "myname",
				Fields: []StructField{
					{
						Name: "field1",
						Type: "fieldType",
						Tags: nil,
					},
					{
						Name: "field2",
						Type: "fieldType",
						Tags: util.NewTags(util.NewTag("tagKey", "tagVal"), util.NewTag("emptyTag", "")),
					},
				},
			}
			want, err := ioutil.ReadFile(filepath.FromSlash("testdata/template_structtype_with_fields.txt"))
			require.NoError(t, err)
			got := execAndFormatTmpl(t, tmpl.Lookup("StructType"), data)
			require.Equal(t, string(want), got)
		})
	})
}

func Test_generateGoFile(t *testing.T) {
	t.Run("CLI", func(t *testing.T) {
		want, err := ioutil.ReadFile(filepath.FromSlash("testdata/generategofile_cli.txt"))
		require.NoError(t, err)

		fileTmpl := FileTmpl{
			CmdHelps: map[string]map[string]string{
				"actions": {
					"cancel-workflow-run": "example",
					"delete-artifact":     "example",
				},
				"activity": {
					"get-feeds":  "example",
					"get-thread": "example",
				},
			},
			FlagHelps: map[string]map[string]map[string]string{
				"actions": {
					"cancel-workflow-run": {
						"owner":  "owner parameter",
						"repo":   "repo parameter",
						"run_id": "run_id parameter",
					},
					"delete-artifact": {
						"artifact_id": "artifact_id parameter",
						"owner":       "owner parameter",
						"repo":        "repo parameter",
					},
				},
				"activity": {
					"get-feeds": {},
					"get-thread": {
						"thread_id": "thread_id parameter",
					},
				},
			},
			PrimaryStructs: []StructTmplHelper{
				{
					Name: "CLI",
					Fields: []StructField{
						{
							Name: "Actions",
							Type: "ActionsCmd",
							Tags: util.NewTags(util.NewTag("cmd", "")),
						},
						{
							Name: "Activity",
							Type: "ActivityCmd",
							Tags: util.NewTags(util.NewTag("cmd", "")),
						},
					},
				},
			},
		}

		got, err := generateGoFile(fileTmpl)
		require.NoError(t, err)
		require.Equal(t, string(want), string(got))
	})

	t.Run("license", func(t *testing.T) {
		want, err := ioutil.ReadFile(filepath.FromSlash("testdata/generategofile_license.txt"))
		require.NoError(t, err)
		fileTmpl := FileTmpl{
			SvcTmpls: []SvcTmpl{
				{
					SvcStruct: StructTmplHelper{
						Name: "LicensesCmd",
						Fields: []StructField{
							{
								Name: "Get",
								Type: "LicensesGetCmd",
								Tags: util.NewTags(util.NewTag("cmd", "")),
							},
							{
								Name: "GetForRepo",
								Type: "LicensesGetForRepoCmd",
								Tags: util.NewTags(util.NewTag("cmd", "")),
							},
							{
								Name: "ListCommonlyUsed",
								Type: "LicensesListCommonlyUsedCmd",
								Tags: util.NewTags(util.NewTag("cmd", "")),
							},
						},
					},
					CmdStructAndMethods: []CmdStructAndMethod{
						{
							CmdStruct: StructTmplHelper{
								Name: "LicensesGetCmd",
								Fields: []StructField{
									{
										Name: "License",
										Type: "string",
										Tags: util.NewTags(util.NewTag("required", ""), util.NewTag("name", "license")),
									},
									{
										Type:   "internal.BaseCmd",
										Import: "github.com/octo-cli/octo-cli/internal",
									},
								},
							},
							RunMethod: RunMethod{
								ReceiverName: "LicensesGetCmd",
								Method:       "GET",
								URLPath:      "/licenses/{license}",
								CodeBlocks: []CodeBlock{
									{
										Code: `
c.UpdateURLPath("license", c.License)`,
									},
								},
							},
						},
						{
							CmdStruct: StructTmplHelper{
								Name: "LicensesGetForRepoCmd",
								Fields: []StructField{
									{
										Name: "Repo",
										Type: "string",
										Tags: util.NewTags(util.NewTag("required", ""), util.NewTag("name", "repo")),
									},
									{
										Type:   "internal.BaseCmd",
										Import: "github.com/octo-cli/octo-cli/internal",
									},
								},
							},
							RunMethod: RunMethod{
								ReceiverName: "LicensesGetForRepoCmd",
								Method:       "GET",
								URLPath:      "/repos/{repo}/license",
								CodeBlocks: []CodeBlock{
									{
										Code: `
c.UpdateURLPath("repo", c.Repo)`,
									},
								},
							},
						},
						{
							CmdStruct: StructTmplHelper{
								Name: "LicensesListCommonlyUsedCmd",
								Fields: []StructField{
									{
										Type:   "internal.BaseCmd",
										Import: "github.com/octo-cli/octo-cli/internal",
									},
								},
							},
							RunMethod: RunMethod{
								ReceiverName: "LicensesListCommonlyUsedCmd",
								Method:       "GET",
								URLPath:      "/licenses",
							},
						},
					},
				},
			},
		}

		got, err := generateGoFile(fileTmpl)
		require.NoError(t, err)
		require.Equal(t, string(want), string(got))
	})
}
