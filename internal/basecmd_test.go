package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_setBodyValue(t *testing.T) {
	body := map[string]interface{}{
		"foo": map[string]interface{}{
			"a": "b",
		},
	}
	setBodyValue(body, []string{"foo", "bar", "baz"}, "qux")
	want := map[string]interface{}{
		"foo": map[string]interface{}{
			"a": "b",
			"bar": map[string]interface{}{
				"baz": "qux",
			},
		},
	}
	require.Equal(t, want, body)
}
