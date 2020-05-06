package tests

import (
	"testing"
)

func TestMarkdown(t *testing.T) {
	t.Run("render", func(t *testing.T) {
		inputText := `
# header

this is the body
`
		want := `<h1>
<a id="user-content-header" class="anchor" href="#header" aria-hidden="true"><span aria-hidden="true" class="octicon octicon-link"></span></a>header</h1>
<p>this is the body</p>

`
		newCmdLine(`markdown`,
			`render`,
			`--text`, inputText,
		).test(t, "test_markdown_render", want, "", false)
	})
}
