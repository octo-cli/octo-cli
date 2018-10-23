package generator

import (
	"fmt"
	"github.com/WillAbides/go-github-cli/generator/internal/routeparser"
	"github.com/fatih/structtag"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var testRtServices map[string]routeparser.Routes

func init() {
	var terr error
	testRtServices, terr = routeparser.ParseRoutesFile("../../testdata/routes.json")
	if terr != nil {
		panic(terr)
	}
}

//func Test_svc_writePkg(t *testing.T) {
//	want, err := ioutil.ReadFile("../../testdata/exampleapp/services/issuessvc/issuessvc.go")
//	require.Nil(t, err)
//	rts := testRtServices["issues"]
//	issueSvc := Svc{
//		Name: "Issues",
//		Commands: []cmd{
//			newCmd("AddLabelsToIssue", rts.FindByIDName("add-labels"), "Owner", "Repo", "Number", "Labels"),
//			newCmd("Create", rts.FindByIDName("create"), "Owner", "Repo"),
//			newCmd("Edit", rts.FindByIDName("edit")),
//			newCmd("List", rts.FindByIDName("list"), "All"),
//			newCmd("Lock", rts.FindByIDName("lock"), "Owner", "Repo", "Number"),
//		},
//	}
//
//	var buf bytes.Buffer
//	err = issueSvc.writePkg(&buf)
//	assert.Nil(t, err)
//	assert.Equal(t, string(want), buf.String())
//}

func nsf(t *testing.T, name, ftype string, tag ...*structtag.Tag) *structField {
	t.Helper()
	tags := newTags(tag...)
	return newStructField(name, ftype, tags)
}

func Test_buildCommandStruct(t *testing.T) {
	t.Run("Issues Edit", func(t *testing.T) {
		rts := testRtServices["issues"]

		s := Svc{Name: "Issues"}
		field, ok := s.getStructField()
		assert.True(t, ok)
		method, ok := field.Type.MethodByName("Edit")
		assert.True(t, ok)
		c := cmd{
			Route:    rts.FindByIDName("edit"),
			ArgNames: []string{"Owner", "Repo", "Number"},
		}
		got, err := buildCommandStruct("Issues", "Edit", method, c)
		assert.Nil(t, err)

		wantRunMethod, err := generateRunMethod("Issues", "Edit", method, "Owner", "Repo", "Number")
		require.Nil(t, err)

		want := &structTmplHelper{
			Name:      "IssuesEditCmd",
			RunMethod: wantRunMethod,
			Fields: []structField{
				*nsf(t, "Token", "string", newTag("env", "GITHUB_TOKEN"), newTag("required", "")),
				*nsf(t, "APIBaseURL", "string", newTag("env", "GITHUB_API_BASE_URL"), newTag("default", "https://api.github.com")),
				*nsf(t, "Owner", "string", newTag("required", "")),
				*nsf(t, "Repo", "string", newTag("required", "")),
				*nsf(t, "Number", "int", newTag("required", "")),
				{Type: "issuesEditCmdIssueRequestFlags"},
			},
		}
		got.OptionsStructs = nil
		assert.Equal(t, want, got)
	})

	t.Run("Issues ListByOrg", func(t *testing.T) {
		rts := testRtServices["issues"]

		s := Svc{Name: "Issues"}
		field, ok := s.getStructField()
		assert.True(t, ok)
		method, ok := field.Type.MethodByName("ListByOrg")
		assert.True(t, ok)
		c := cmd{
			Route:    rts.FindByIDName("list-for-org"),
			ArgNames: []string{"Org"},
		}
		got, err := buildCommandStruct("Issues", "ListByOrg", method, c)
		assert.Nil(t, err)

		want := &structTmplHelper{
			Name: "IssuesListByOrgCmd",
			Fields: []structField{
				*nsf(t, "Token", "string", newTag("env", "GITHUB_TOKEN"), newTag("required", "")),
				*nsf(t, "APIBaseURL", "string", newTag("env", "GITHUB_API_BASE_URL"), newTag("default", "https://api.github.com")),
				*nsf(t, "Org", "string", newTag("required", "")),
				{Type: "issuesListByOrgCmdIssueListOptionsFlags"},
			},
		}

		assert.Equal(t, want.Name, got.Name)
		assert.Equal(t, want.Fields, got.Fields)
	})

}

func Test_generateRunMethod(t *testing.T) {
	t.Run("Issues Edit", func(t *testing.T) {
		want := &runMethodHelper{
			StructName: "IssuesEditCmd",
			HasElement: true,
			SvcName:    "Issues",
			FuncName:   "Edit",
			Args: []runMethodArgHelper{
				{Name: "Owner"},
				{Name: "Repo"},
				{Name: "Number"},
				{Name: "IssueRequest", IsPtr: true},
			},
		}

		s := Svc{Name: "Issues"}
		field, ok := s.getStructField()
		require.True(t, ok)
		method, ok := field.Type.MethodByName("Edit")
		assert.True(t, ok)
		got, err := generateRunMethod("Issues", "Edit", method, "Owner", "Repo", "Number")
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("Issues Lock", func(t *testing.T) {
		want := &runMethodHelper{
			StructName: "IssuesLockCmd",
			HasElement: false,
			SvcName:    "Issues",
			FuncName:   "Lock",
			Args: []runMethodArgHelper{
				{Name: "Owner"},
				{Name: "Repo"},
				{Name: "Number"},
				{Name: "LockIssueOptions", IsPtr: true},
			},
		}

		s := Svc{Name: "Issues"}
		field, ok := s.getStructField()
		assert.True(t, ok)
		method, ok := field.Type.MethodByName("Lock")
		assert.True(t, ok)
		got, err := generateRunMethod("Issues", "Lock", method, "Owner", "Repo", "Number")
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	})
}
func Test_fieldFlagName(t *testing.T) {
	type example struct {
		Body       *string  `json:"body,omitempty"`
		Labels     []string `url:"labels,comma,omitempty"`
		LockReason string   `json:"lock_reason,omitempty"`
		NoTag      string
		JSONDiff   *string `json:"something_different,omitempty"`
	}

	t.Run("no tag", func(t *testing.T) {
		field, _ := reflect.TypeOf(example{}).FieldByName("NoTag")
		assert.Equal(t, "no-tag", fieldFlagName(field))
	})

	t.Run("from json", func(t *testing.T) {
		field, _ := reflect.TypeOf(example{}).FieldByName("JSONDiff")
		assert.Equal(t, "something-different", fieldFlagName(field))
	})

	t.Run("non-json tag", func(t *testing.T) {
		field, _ := reflect.TypeOf(example{}).FieldByName("Labels")
		assert.Equal(t, "labels", fieldFlagName(field))
	})
}

func Test_getStructFields(t *testing.T) {
	t.Run("github.ListOptions", func(t *testing.T) {
		fields := typeToFields(reflect.TypeOf(github.ListOptions{}))
		want := []structField{
			*nsf(t, "Page", "int", newTag("name", "page")),
			*nsf(t, "PerPage", "int", newTag("name", "per-page")),
		}
		got := getStructFields(fields, nil)
		assert.Equal(t, want, got)
	})

	t.Run("github.IssueListOptions", func(t *testing.T) {
		fields := typeToFields(reflect.TypeOf(github.IssueListOptions{}))
		want := []structField{
			*nsf(t, "Filter", "string", newTag("name", "filter")),
			*nsf(t, "State", "string", newTag("name", "state")),
			*nsf(t, "Labels", "[]string", newTag("name", "labels")),
			*nsf(t, "Sort", "string", newTag("name", "sort")),
			*nsf(t, "Direction", "string", newTag("name", "direction")),
			*nsf(t, "Since", "time.Time", newTag("name", "since")),
			*nsf(t, "Page", "int", newTag("name", "page")),
			*nsf(t, "PerPage", "int", newTag("name", "per-page")),
		}
		got := getStructFields(fields, nil)
		assert.Equal(t, want, got)
	})
}

func Test_generateOptionsStruct(t *testing.T) {
	t.Run("fields", func(t *testing.T) {
		t.Run("github.ListOptions", func(t *testing.T) {
			mt := reflect.TypeOf(github.ListOptions{})
			oStruct := generateOptionsStruct("", mt, nil)
			want := []structField{
				*nsf(t, "Page", "int", newTag("name", "page")),
				*nsf(t, "PerPage", "int", newTag("name", "per-page")),
			}
			assert.Equal(t, want, oStruct.MainStruct.Fields)
			fmt.Println(oStruct.ToFunc.ValSetters)
		})

		t.Run("github.IssueListOptions", func(t *testing.T) {
			mt := reflect.TypeOf(github.IssueListOptions{})
			oStruct := generateOptionsStruct("", mt, nil)
			want := []structField{
				*nsf(t, "Filter", "string", newTag("name", "filter")),
				*nsf(t, "State", "string", newTag("name", "state")),
				*nsf(t, "Labels", "[]string", newTag("name", "labels")),
				*nsf(t, "Sort", "string", newTag("name", "sort")),
				*nsf(t, "Direction", "string", newTag("name", "direction")),
				*nsf(t, "Since", "time.Time", newTag("name", "since")),
				*nsf(t, "Page", "int", newTag("name", "page")),
				*nsf(t, "PerPage", "int", newTag("name", "per-page")),
			}
			assert.Equal(t, want, oStruct.MainStruct.Fields)
			fmt.Println(oStruct.ToFunc.ValSetters)
		})
	})

	type anonStruct struct {
		AnonOne string `json:"anon1,omitempty"`
	}

	type anonPtrStruct struct {
		AnonPtrOne string `json:"anonPtr1,omitempty"`
	}

	type example struct {
		Body       *string  `json:"body,omitempty"`
		Labels     []string `url:"labels,comma,omitempty"`
		LockReason string   `json:"lock_reason,omitempty"`
		NoTag      string
		JSONDiff   *string `json:"something_different,omitempty"`
		anonStruct     //nolint:megacheck
		*anonPtrStruct //nolint:megacheck
	}

	t.Run("val setters", func(t *testing.T) {
		t.Run("ex", func(t *testing.T) {
			mt := reflect.TypeOf(example{})
			oStruct := generateOptionsStruct("", mt, nil)
			want := []valSetter{
				{TargetIsPtr: true, Name: "Body", FlagName: "body"},
				{TargetIsPtr: false, Name: "Labels", FlagName: "labels"},
				{TargetIsPtr: false, Name: "LockReason", FlagName: "lock-reason"},
				{TargetIsPtr: false, Name: "NoTag", FlagName: "no-tag"},
				{TargetIsPtr: true, Name: "JSONDiff", FlagName: "something-different"},
				{TargetIsPtr: false, Name: "AnonOne", FlagName: "anon1"},
				{TargetIsPtr: false, Name: "AnonPtrOne", FlagName: "anonptr1"},
			}
			assert.Equal(t, want, oStruct.ToFunc.ValSetters)
		})

		t.Run("github.IssueListOptions", func(t *testing.T) {
			mt := reflect.TypeOf(github.IssueListOptions{})
			oStruct := generateOptionsStruct("", mt, nil)

			want := []valSetter{
				{Name: "Filter", FlagName: "filter", TargetIsPtr: false},
				{Name: "State", FlagName: "state", TargetIsPtr: false},
				{Name: "Labels", FlagName: "labels", TargetIsPtr: false},
				{Name: "Sort", FlagName: "sort", TargetIsPtr: false},
				{Name: "Direction", FlagName: "direction", TargetIsPtr: false},
				{Name: "Since", FlagName: "since", TargetIsPtr: false},
				{Name: "Page", FlagName: "page", TargetIsPtr: false},
				{Name: "PerPage", FlagName: "per-page", TargetIsPtr: false},
			}

			assert.Equal(t, want, oStruct.ToFunc.ValSetters)
		})

		t.Run("github.ListOptions", func(t *testing.T) {
			mt := reflect.TypeOf(github.ListOptions{})
			oStruct := generateOptionsStruct("", mt, nil)

			want := []valSetter{
				{Name: "Page", FlagName: "page", TargetIsPtr: false},
				{Name: "PerPage", FlagName: "per-page", TargetIsPtr: false},
			}

			assert.Equal(t, want, oStruct.ToFunc.ValSetters)
		})
	})

}
