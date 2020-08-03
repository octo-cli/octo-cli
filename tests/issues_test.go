package tests

import (
	"testing"
)

func TestIssues(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		format := `{{.state}} : {{.title}} : {{.body}}`
		newCmdLine(`issues`,
			`create`,
			`--repo=octo-cli-testorg/test-create-issue`,
			`--title=test this`,
			`--body=test this body`,
			`--labels=label1`,
			`--labels=label2`,
			`--milestone=1`,
			`--assignees=octo-cli-testuser`,
			`--format`, format).
			test(t, "test_issues_create", "open : test this : test this body\n", "", false)

	})

	t.Run("get", func(t *testing.T) {
		format := `{{.state}} : {{.title}} : {{.body}}`
		newCmdLine(`issues`,
			`get`,
			`--repo=octo-cli-testorg/test-create-issue`,
			`--issue_number=1`,
			`--format`, format,
		).test(t, "test_issues_get", "open : \"test this\" : \"test this body\"\n", "", false)
	})

	t.Run("add-labels", func(t *testing.T) {
		format := `{{range .}}{{.name}} {{end}}`
		newCmdLine(`issues`,
			`add-labels`,
			`--repo=octo-cli-testorg/test-create-issue`,
			`--issue_number=1`,
			`--labels`, `foo`,
			`--labels=bar`,
			`--format`, format,
		).test(t, "test_issues_add-labels", "bar foo label1 label2 \n", "", false)
	})

	t.Run("remove-label", func(t *testing.T) {
		format := `{{.name}} `
		newCmdLine(`issues`,
			`remove-label`,
			`--repo=octo-cli-testorg/test-create-issue`,
			`--issue_number=1`,
			`--name`, `foo`,
			`--format`, format,
			`--output-each`, `.`,
		).test(t, "test_issues_remove-label", "bar label1 label2 \n", "", false)
	})
}
