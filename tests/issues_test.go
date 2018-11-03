package tests

import (
	"encoding/json"
	"github.com/octo-cli/octo-cli/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssues(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		stdout, stderr, err := testCmdLine(t, "test_issues_create", &services.IssuesCmd{},
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
		assert.Equal(t, "test this", got["title"])
		assert.Equal(t, "test this body", got["body"])
		assert.Len(t, got["labels"], 2)
		assert.EqualValues(t, 1, got["milestone"].(map[string]interface{})["number"])
	})
}
