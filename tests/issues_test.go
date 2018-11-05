package tests

import (
	"encoding/json"
	"fmt"
	"github.com/octo-cli/octo-cli/internal/generated"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"text/tabwriter"
	"text/template"
)

func TestIssues(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		format := `{{.state}} : {{.title}} : {{.body}}`
		newCmdLine(`issues`,
			`create`,
			`--owner=octo-cli-testorg`,
			`--repo=test-create-issue`,
			`--title=test this`,
			`--body=test this body`,
			`--labels=label1`,
			`--labels=label2`,
			`--milestone=1`,
			`--assignees=octo-cli-testuser`,
			`--format`, format).
			test(t, "test_issues_create", "open : test this : test this body\n", "", false)

	})

	t.Run("Get", func(t *testing.T) {
		format := `{{.state}} : {{.title}} : {{.body}}`
		newCmdLine(`issues`,
			`get`,
			`--owner=octo-cli-testorg`,
			`--repo=test-create-issue`,
			`--number=1`,
			`--format`, format,
		).test(t, "test_issues_get",  "open : \"test this\" : \"test this body\"\n", "", false)
	})

	t.Run("play with  output", func(t *testing.T) {
		// just keeping this here until I put it somewhere else because I don't want to forget how it works.
		stdout, stderr, err := runCmdLine(t, "test_issues_create", &generated.IssuesCmd{},
			`create`,
			`--owner=octo-cli-testorg`,
			`--repo=test-create-issue`,
			`--title=test this`,
			`--body=test this body`,
			`--labels=label1`,
			`--labels=label2`,
			`--milestone=1`,
			`--assignees=octo-cli-testuser`,
			`--raw-output`,
		)
		assert.NoError(t, err)
		assert.Empty(t, stderr.String())
		var got map[string]interface{}
		err = json.Unmarshal(stdout.Bytes(), &got)
		assert.NoError(t, err)

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
		assert.NoError(t, err)
		w := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)
		err = tp.Execute(w, got)
		assert.NoError(t, err)
		err = w.Flush()
		assert.NoError(t, err)

	})
}
