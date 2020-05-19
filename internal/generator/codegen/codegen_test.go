package codegen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateGoFile(t *testing.T) {
	t.Run("CLI", func(t *testing.T) {
		want, err := ioutil.ReadFile(filepath.FromSlash("testdata/generategofile_cli.txt"))
		require.NoError(t, err)

		fileTmpl1 := fileTmpl{
			cmdHelps: map[string]map[string]string{
				"actions": {
					"cancel-workflow-run": "example",
					"delete-artifact":     "example",
				},
				"activity": {
					"get-feeds":  "example",
					"get-thread": "example",
				},
			},
			flagHelps: map[string]map[string]map[string]string{
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
			primaryStructs: []structTmplHelper{
				{
					name: "CLI",
					fields: []structField{
						{
							name:      "Actions",
							fieldType: "ActionsCmd",
							tags:      map[string]string{"cmd": ""},
						},
						{
							name:      "Activity",
							fieldType: "ActivityCmd",
							tags:      map[string]string{"cmd": ""},
						},
					},
				},
			},
		}

		got, err := generateGoFile(fileTmpl1)
		require.NoError(t, err)
		require.Equal(t, string(want), string(got))
	})

	t.Run("license", func(t *testing.T) {
		want, err := ioutil.ReadFile(filepath.FromSlash("testdata/generategofile_license.txt"))
		require.NoError(t, err)
		fileTmpl1 := fileTmpl{
			svcTmpls: []svcTmpl{
				{
					svcStruct: &structTmplHelper{
						name: "LicensesCmd",
						fields: []structField{
							{
								name:      "Get",
								fieldType: "LicensesGetCmd",
								tags:      map[string]string{"cmd": ""},
							},
							{
								name:      "GetForRepo",
								fieldType: "LicensesGetForRepoCmd",
								tags:      map[string]string{"cmd": ""},
							},
							{
								name:      "ListCommonlyUsed",
								fieldType: "LicensesListCommonlyUsedCmd",
								tags:      map[string]string{"cmd": ""},
							},
						},
					},
					cmdStructAndMethods: []cmdStructAndMethod{
						{
							cmdStruct: &structTmplHelper{
								name: "LicensesGetCmd",
								fields: []structField{
									{
										name:      "License",
										fieldType: "string",
										tags:      map[string]string{"required": "true", "name": "license"}},
									{
										fieldType: "internal.BaseCmd",
									},
								},
							},
							runMethod: &runMethod{
								receiverName:    "LicensesGetCmd",
								method:          "GET",
								urlPath:         "/licenses/{license}",
								codeGroupAdders: []codeGroupAdder{newRunMethodAdder("license", "c.UpdateURLPath")},
							},
						},
						{
							cmdStruct: &structTmplHelper{
								name: "LicensesGetForRepoCmd",
								fields: []structField{
									{
										name:      "Repo",
										fieldType: "string",
										tags:      map[string]string{"required": "true", "name": "repo"}},
									{
										fieldType: "internal.BaseCmd",
									},
								},
							},
							runMethod: &runMethod{
								receiverName:    "LicensesGetForRepoCmd",
								method:          "GET",
								urlPath:         "/repos/{repo}/license",
								codeGroupAdders: []codeGroupAdder{newRunMethodAdder("repo", "c.UpdateURLPath")},
							},
						},
						{
							cmdStruct: &structTmplHelper{
								name: "LicensesListCommonlyUsedCmd",
								fields: []structField{
									{
										fieldType: "internal.BaseCmd",
									},
								},
							},
							runMethod: &runMethod{
								receiverName: "LicensesListCommonlyUsedCmd",
								method:       "GET",
								urlPath:      "/licenses",
							},
						},
					},
				},
			},
		}

		got, err := generateGoFile(fileTmpl1)
		require.NoError(t, err)
		require.Equal(t, string(want), string(got))
	})
}
