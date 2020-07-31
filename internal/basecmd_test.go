package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestBaseCmd_OutputResult(t *testing.T) {
	type args struct {
		respBody   string
		rawOutput  bool
		format     string
		outputEach string
	}
	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantErr    bool
	}{
		{
			name: "raw text",
			args: args{
				respBody:  "hello",
				rawOutput: true,
				format:    "",
			},
			wantStdout: "hello",
			wantErr:    false,
		},
		{
			name: "raw json",
			args: args{
				respBody:  `{"hello": "hi"}`,
				rawOutput: true,
				format:    "",
			},
			wantStdout: `{"hello": "hi"}`,
			wantErr:    false,
		},
		{
			name: "pretty json",
			args: args{
				respBody:  `{"hello": "hi"}`,
				rawOutput: false,
				format:    "",
			},
			wantStdout: `{
  "hello": "hi"
}`,
			wantErr: false,
		},
		{
			name: "invalid json",
			args: args{
				respBody:  `{"hello": "hi"`,
				rawOutput: false,
				format:    "",
			},
			wantStdout: "",
			wantErr:    true,
		},
		{
			name: "format",
			args: args{
				respBody:  `{"hello": "hi"}`,
				rawOutput: false,
				format:    "hello {{.hello}}",
			},
			wantStdout: "hello hi",
			wantErr:    false,
		},
		{
			name: "output-each",
			args: args{
				respBody: `{
  "foo": {
    "bar": [
      {
        "baz": 1,
        "qux": true
      },
      {
        "baz": 2,
        "qux": true
      },
      {
        "qux": true
      }
    ]
  }
}
`,
				format:     "{{.baz}},",
				outputEach: ".foo.bar",
			},
			wantStdout: `1,2,<no value>,`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			resp := &http.Response{
				Body:   ioutil.NopCloser(strings.NewReader(tt.args.respBody)),
				Header: http.Header{},
			}
			resp.Header.Set("content-type", "application/json")
			cmd := &BaseCmd{
				RawOutput:  tt.args.rawOutput,
				Format:     tt.args.format,
				OutputEach: tt.args.outputEach,
			}
			if err := cmd.OutputResult(resp, stdout); (err != nil) != tt.wantErr {
				t.Errorf("OutputResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantStdout, strings.TrimSuffix(stdout.String(), "\n"))

		})
	}
}
