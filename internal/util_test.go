package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrettyPrintJson(t *testing.T) {
	for _, test := range []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "single object",
			input: `{ "foo": {"bar": ["baz", "qux"]} }`,
			want: `{
  "foo": {
    "bar": [
      "baz",
      "qux"
    ]
  }
}`,
		},
		{
			name:  "string array",
			input: `[ "foo", "bar" ]`,
			want: `[
  "foo",
  "bar"
]`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := prettyPrintJSON([]byte(test.input))
			assert.Nil(t, err)
			assert.Equal(t, test.want, string(got))

		})
	}
}
