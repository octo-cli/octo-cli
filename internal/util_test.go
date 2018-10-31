package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prettyPrintJSON(t *testing.T) {
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

func TestOutputResult(t *testing.T) {
	type args struct {
		respBody string
		rawJSON  bool
		format   string
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
				respBody: "hello",
				rawJSON:  true,
				format:   "",
			},
			wantStdout: "hello",
			wantErr:    false,
		},
		{
			name: "raw json",
			args: args{
				respBody: `{"hello": "hi"}`,
				rawJSON:  true,
				format:   "",
			},
			wantStdout: `{"hello": "hi"}`,
			wantErr:    false,
		},
		{
			name: "pretty json",
			args: args{
				respBody: `{"hello": "hi"}`,
				rawJSON:  false,
				format:   "",
			},
			wantStdout: `{
  "hello": "hi"
}`,
			wantErr: false,
		},
		{
			name: "invalid json",
			args: args{
				respBody: `{"hello": "hi"`,
				rawJSON:  false,
				format:   "",
			},
			wantStdout: "",
			wantErr:    true,
		},
		{
			name: "format",
			args: args{
				respBody: `{"hello": "hi"}`,
				rawJSON:  false,
				format:   "hello {{.hello}}",
			},
			wantStdout: "hello hi",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			resp := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(tt.args.respBody)),
			}
			if err := OutputResult(resp, tt.args.rawJSON, tt.args.format, stdout); (err != nil) != tt.wantErr {
				t.Errorf("OutputResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantStdout, stdout.String())

		})
	}
}
