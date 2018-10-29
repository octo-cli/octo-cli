package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"text/tabwriter"
	"text/template"

	"github.com/alecthomas/kong"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"
)

func init() {
	tkn, ok := os.LookupEnv("TESTUSER_TOKEN")
	if !ok {
		tkn = "deadbeef"
	}
	_ = os.Setenv("GITHUB_TOKEN", tkn)
}

func startVCR(t *testing.T, recPath string) *recorder.Recorder {
	t.Helper()
	var err error
	rec, err := recorder.New(recPath)
	require.Nil(t, err)
	customTransport = rec
	return rec
}

func testCmdLine(t *testing.T, fixtureName string, cmdStruct interface{}, cmdline ...string) (sout bytes.Buffer, serr bytes.Buffer, err error) {
	t.Helper()
	rec := startVCR(t, filepath.Join("testdata", "fixtures", fixtureName))
	defer func() {_ = rec.Stop()}()
	p, e := kong.New(cmdStruct)

	require.Nil(t, e)
	stdout = &sout
	stderr = &serr

	k, e := p.Parse(cmdline)
	require.Nil(t, e)
	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	err = k.Run(valueIsSetMap)
	return
}

func TestCreate(t *testing.T) {
	stdout, stderr, err := testCmdLine(t, "test_issues_create", &IssuesCmd{},
		`create`,
		`--owner=go-github-cli-testorg`,
		`--repo=test-create-issue`,
		`--title=test this`,
		`--body=test this body`,
		`--labels=label1`,
		`--labels=label2`,
		`--milestone=1`,
		`--assignees=go-github-cli-testuser`,
	)
	assert.Nil(t, err)
	assert.Empty(t, stderr.String())
	var got map[string]interface{}
	err = json.Unmarshal(stdout.Bytes(), &got)
	assert.Nil(t, err)
	assert.Equal(t, "test this", got["title"])
	assert.Equal(t, "test this body", got["body"])
	assert.Len(t, got["labels"], 2)
	assert.EqualValues(t, 1, got["milestone"].(map[string]interface{})["number"])
}

// just keeping this here until I put it somewhere else because I don't want to forget how it works.
func TestPlayWithOutput(t *testing.T) {
	stdout, stderr, err := testCmdLine(t, "test_issues_create", &IssuesCmd{},
		`create`,
		`--owner=go-github-cli-testorg`,
		`--repo=test-create-issue`,
		`--title=test this`,
		`--body=test this body`,
		`--labels=label1`,
		`--labels=label2`,
		`--milestone=1`,
		`--assignees=go-github-cli-testuser`,
	)
	assert.Nil(t, err)
	assert.Empty(t, stderr.String())
	var got map[string]interface{}
	err = json.Unmarshal(stdout.Bytes(), &got)
	assert.Nil(t, err)

	for k, _ := range got {
		fmt.Println(k)
	}

	tp, err := template.New("").Funcs(template.FuncMap{"wider": func(a, b interface{}) string { return fmt.Sprintf("%-20s%d", a, b) }}).Parse(`
	Number	{{.number}}
	State	{{.state}}
	Title	{{.title}}
	User	{{.user.login}}
	URL	{{.html_url}}
	CreatedAt	{{.created_at}}
	UpdatedAt	{{.updated_at}}
	{{if .labels}}Labels{{range .labels}}	{{.name}}
	{{end}}{{end}}{{if .milestone}}Milestone	{{.milestone.title}}{{end}}
	{{if .assignees}}Assignees{{range .assignees}}	{{.login}}
	{{end}}{{end}}
	`)
	assert.Nil(t, err)
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)
	err = tp.Execute(w, got)
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
}
