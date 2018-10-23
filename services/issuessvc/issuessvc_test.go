package issuessvc

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"text/tabwriter"
	"text/template"
)

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
	gotIssue := github.Issue{}
	err = json.Unmarshal(stdout.Bytes(), &gotIssue)
	assert.Nil(t, err)
	assert.Equal(t, "test this", gotIssue.GetTitle())
	assert.Equal(t, "test this body", gotIssue.GetBody())
	assert.Len(t, gotIssue.Labels, 2)
	assert.Equal(t, 1, gotIssue.GetMilestone().GetNumber())
	assert.Len(t, gotIssue.Assignees, 1)
	assert.Equal(t, "go-github-cli-testuser", *gotIssue.Assignees[0].Login)
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
	gotIssue := github.Issue{}
	err = json.Unmarshal(stdout.Bytes(), &gotIssue)
	assert.Nil(t, err)

	tp, err := template.New("").Funcs(template.FuncMap{"wider": func(a, b interface{}) string { return fmt.Sprintf("%-20s%d", a, b) }}).Parse(`
	Number	{{.Number}}
	State	{{.State}}
	Title	{{.Title}}
	User	{{.User.Login}}
	URL	{{.HTMLURL}}
	CreatedAt	{{.CreatedAt}}
	UpdatedAt	{{.UpdatedAt}}
	{{if .Labels}}Labels{{range .Labels}}	{{.Name}}
	{{end}}{{end}}{{if .Milestone}}Milestone	{{.Milestone.Title}}{{end}}
	{{if .Assignees}}Assignees{{range .Assignees}}	{{.Login}}
	{{end}}{{end}}
	`)
	assert.Nil(t, err)

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)
	err = tp.Execute(w, gotIssue)
	assert.Nil(t, err)
	err = w.Flush()
	assert.Nil(t, err)
}
