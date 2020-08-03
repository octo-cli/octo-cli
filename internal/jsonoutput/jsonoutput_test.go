package jsonoutput

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	jsonNull   = `null`
	jsonString = `"jsonString"`
	jsonZero   = `0`
	jsonTwelve = `12`
	jsonMap    = `{"one": 1, "two": 2, "three": 3}`
	jsonList   = `[1, "two"]`
)

func assertFormatJSONOutput(t *testing.T, jsonData, format, want string) {
	t.Helper()
	got, err := FormatJSONOutput([]byte(jsonData), format, "")
	require.NoError(t, err)
	require.Equal(t, want, string(got))
}

func Test_FormatJSONOutput(t *testing.T) {
	t.Run("invalid json", func(t *testing.T) {
		got, err := FormatJSONOutput([]byte("invalid json"), "{{ . }}", "")
		require.EqualError(t, err, "invalid json")
		require.Empty(t, got)
	})

	t.Run("invalid template", func(t *testing.T) {
		got, err := FormatJSONOutput([]byte(jsonMap), "{{ . ", "")
		require.EqualError(t, err, "failed parsing format template")
		require.Empty(t, got)
	})

	t.Run("exec template error", func(t *testing.T) {
		got, err := FormatJSONOutput([]byte(jsonMap), "{{ .one.foo.bar }}", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "error executing template")
		require.Empty(t, got)
	})

	t.Run("outputEach", func(t *testing.T) {
		data := fmt.Sprintf(`{"foo": [%s, %s, %s]}`, jsonMap, jsonMap, jsonMap)
		format := `{{.one | newLine }}`
		got, err := FormatJSONOutput([]byte(data), format, ".foo")
		require.NoError(t, err)
		require.Equal(t, `1
1
1
`, string(got))
	})
}

func Test_FormatJSONOutput_funcs(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonNull, `{{default "foo" .}}`, "foo")
		assertFormatJSONOutput(t, jsonString, `{{default "foo" .}}`, "jsonString")
	})

	t.Run("empty", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonNull, `{{empty .}}`, "true")
		assertFormatJSONOutput(t, jsonString, `{{empty .}}`, "false")
		assertFormatJSONOutput(t, jsonZero, `{{empty .}}`, "true")
	})

	t.Run("ternary", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonNull, `{{ternary "foo" "bar" true}}`, "foo")
		assertFormatJSONOutput(t, jsonNull, `{{ternary "foo" "bar" false}}`, "bar")
		assertFormatJSONOutput(t, jsonNull, `{{ternary "foo" "bar" (not (empty .))}}`, "bar")
	})

	t.Run("coalesce", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonNull, `{{coalesce .}}`, `<no value>`)
		assertFormatJSONOutput(t, jsonNull, `{{coalesce 0 0 0}}`, `<no value>`)
		assertFormatJSONOutput(t, jsonNull, `{{coalesce 0 10 20}}`, `10`)
	})

	t.Run("trim", func(t *testing.T) {
		assertFormatJSONOutput(t, `"\n\t\tfoo\t\t\n"`, `{{trim .}}`, `foo`)
	})

	t.Run("trimAll", func(t *testing.T) {
		assertFormatJSONOutput(t, `"abc foo abc"`, `{{. | trimAll "abc"}}`, ` foo `)
	})

	t.Run("trimSuffix", func(t *testing.T) {
		assertFormatJSONOutput(t, `"abc foo abc"`, `{{. | trimSuffix "abc"}}`, `abc foo `)
	})

	t.Run("trimPrefix", func(t *testing.T) {
		assertFormatJSONOutput(t, `"abc foo abc"`, `{{. | trimPrefix "abc"}}`, ` foo abc`)
	})

	t.Run("upper", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{upper .}}`, "JSONSTRING")
	})

	t.Run("lower", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{lower .}}`, "jsonstring")
	})

	t.Run("substr", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{. | substr 4 7 }}`, "Str")
		assertFormatJSONOutput(t, jsonString, `{{. | substr 4 17 }}`, "String")
	})

	t.Run("trunc", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{. | trunc 4 }}`, "json")
		assertFormatJSONOutput(t, jsonString, `{{. | trunc -3 }}`, "ing")
	})

	t.Run("contains", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{. | contains "json" }}`, "true")
		assertFormatJSONOutput(t, jsonString, `{{. | contains "bar" }}`, "false")
	})

	t.Run("replace", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{. | replace "String" "Thing" }}`, "jsonThing")
	})

	t.Run("cat", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{ cat . "foo" }}`, "jsonString foo")
	})

	t.Run("split", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonString, `{{ . | split "St" }}`, "[json ring]")
		assertFormatJSONOutput(t, jsonString, `{{ . | split "" }}`, "[j s o n S t r i n g]")
	})

	t.Run("sortAlpha", func(t *testing.T) {
		assertFormatJSONOutput(t, `["foo", "bar", "baz"]`, `{{ . | sortAlpha }}`, "[bar baz foo]")
	})

	t.Run("toString", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonTwelve, `{{ . | toString | printf "%T" }}`, "string")
		assertFormatJSONOutput(t, jsonTwelve, `{{ . | toString  }}`, "12")
	})

	t.Run("toStrings", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonTwelve, `{{ . | toStrings | printf "%T" }}`, "[]string")
		assertFormatJSONOutput(t, jsonTwelve, `{{ . | toStrings }}`, "[12]")
	})

	t.Run("base", func(t *testing.T) {
		assertFormatJSONOutput(t, `"foo/bar/baz.txt"`, `{{ . | base }}`, "baz.txt")
	})

	t.Run("dir", func(t *testing.T) {
		assertFormatJSONOutput(t, `"foo/bar/baz.txt"`, `{{ . | dir }}`, "foo/bar")
	})

	t.Run("clean", func(t *testing.T) {
		assertFormatJSONOutput(t, `"foo/bar/../baz.txt"`, `{{ . | clean }}`, "foo/baz.txt")
	})

	t.Run("ext", func(t *testing.T) {
		assertFormatJSONOutput(t, `"foo/bar/baz.txt"`, `{{ . | ext }}`, ".txt")
	})

	t.Run("keys", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonMap, `{{ keys . | sortAlpha }}`, "[one three two]")
	})

	t.Run("pick", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonMap, `{{ pick . "one" "two" "four" | keys | sortAlpha }}`, "[one two]")
	})

	t.Run("omit", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonMap, `{{ omit . "one" "two" "four" | keys | sortAlpha }}`, "[three]")
	})

	t.Run("get", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonMap, `{{ get . "one" }}`, "1")
	})

	t.Run("hasKey", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonMap, `{{ hasKey . "one" }}`, "true")
		assertFormatJSONOutput(t, jsonMap, `{{ hasKey . "four" }}`, "false")
	})

	t.Run("list", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonNull, `{{ list 1 "two" }}`, "[1 two]")
	})

	t.Run("toJson", func(t *testing.T) {
		want := `{"one":1,"three":3,"two":2}`
		assertFormatJSONOutput(t, jsonMap, `{{ . | toJson }}`, want)
	})

	t.Run("toRawJson", func(t *testing.T) {
		want := `{"one":1,"three":3,"two":2}`
		assertFormatJSONOutput(t, jsonMap, `{{ . | toRawJson }}`, want)
	})

	t.Run("toPrettyJson", func(t *testing.T) {
		want := `{
  "one": 1,
  "three": 3,
  "two": 2
}`
		assertFormatJSONOutput(t, jsonMap, `{{ . | toPrettyJson }}`, want)
	})

	t.Run("first", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonList, `{{ . | first }}`, "1")
	})

	t.Run("last", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonList, `{{ . | last }}`, "two")
	})

	t.Run("has", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonList, `{{ . | has "two" }}`, "true")
		assertFormatJSONOutput(t, jsonList, `{{ . | has "foo" }}`, "false")
	})

	t.Run("compact", func(t *testing.T) {
		data := `[1, 0, null, {}, [], "two"]`
		assertFormatJSONOutput(t, data, `{{ . | compact }}`, "[1 two]")
	})

	t.Run("slice", func(t *testing.T) {
		data := `[1,2,3,4,5]`
		assertFormatJSONOutput(t, data, `{{ slice . 1 4 }}`, "[2 3 4]")
	})
}

func Test_FormatJSONOutput_newLine(t *testing.T) {
	assertFormatJSONOutput(t, jsonString, `{{. | newLine}}`, "jsonString\n")
	assertFormatJSONOutput(t, jsonNull, `{{. | newLine}}`, "<nil>\n")
	assertFormatJSONOutput(t, jsonNull, `{{"" | newLine}}`, "\n")
	assertFormatJSONOutput(t, jsonTwelve, `{{. | newLine}}`, "12\n")
	assertFormatJSONOutput(t, jsonList, `{{. | newLine}}`, "[1 two]\n")
}

func Test_FormatJSONOutput_toCsv(t *testing.T) {
	t.Run("list", func(t *testing.T) {
		data := `["foo", 2, [1,2,3], "a,b,\"c\"", "hello\tworld"]`
		want := `foo,2,[1 2 3],"a,b,""c""",hello\tworld` + "\n"
		assertFormatJSONOutput(t, data, `{{. | toCsv}}`, want)
	})

	t.Run("vals", func(t *testing.T) {
		assertFormatJSONOutput(t, jsonList, `{{toCsv "a" 12 "c" .}}`, `a,12,c,[1 two]`+"\n")
	})
}

func Test_FormatJSONOutput_toTsv(t *testing.T) {
	data := `["foo", 2, [1,2,3], "a,b,\"c\"", "hello\tworld"]`
	want := `foo	2	[1 2 3]	"a,b,""c"""	hello\tworld` + "\n"
	assertFormatJSONOutput(t, data, `{{. | toTsv}}`, want)
}

func Test_FormatJSONOutput_fromBase64(t *testing.T) {
	data := `"Zm9vCmJhcgo="` // base64 of "foo\nbar\n"
	assertFormatJSONOutput(t, data, `{{. | fromBase64}}`, "foo\nbar\n")
}

func Test_FormatJSONOutput_join(t *testing.T) {
	assertFormatJSONOutput(t, jsonList, `{{join "," .}}`, "1,two")
	assertFormatJSONOutput(t, jsonList, `{{join "," . .}}`, "[1 two],[1 two]")
	assertFormatJSONOutput(t, jsonNull, `{{join "," .}}`, "")
	assertFormatJSONOutput(t, `[]`, `{{join "," .}}`, "")
	assertFormatJSONOutput(t, jsonMap, `{{join "," .one .two .three}}`, "1,2,3")
	assertFormatJSONOutput(t, jsonNull, `{{join ","}}`, "")
}

func Test_FormatJSONOutput_toYaml(t *testing.T) {
	assertFormatJSONOutput(t, jsonNull, `{{. | toYaml}}`, "null\n")
	assertFormatJSONOutput(t, jsonMap, `{{. | toYaml}}`, `one: 1
three: 3
two: 2
`)

	assertFormatJSONOutput(t, jsonList, `{{. | toYaml}}`, `- 1
- two
`)
}

func Test_FormatJSONOutput_pluck(t *testing.T) {
	assertFormatJSONOutput(t, jsonNull, `{{pluck "foo" }}`, "[]")
	assertFormatJSONOutput(t, jsonMap, `{{pluck "one" . }}`, "[1]")
	assertFormatJSONOutput(t, jsonMap, `{{pluck "four" . }}`, "[]")
	mapArray := fmt.Sprintf(`[%s, %s]`, jsonMap, jsonMap)
	assertFormatJSONOutput(t, mapArray, `{{pluck "one" . }}`, "[1 1]")
	assertFormatJSONOutput(t, mapArray, `{{pluck "four" . }}`, "[]")
	assertFormatJSONOutput(t, jsonMap, `{{pluck "one" . . "foo" .}}`, "[1 1 1]")
	assertFormatJSONOutput(t, jsonMap, `{{pluck "four" . . .}}`, "[]")
}

func Test_fromBase64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		data := "Zm9vCmJhcgo=" // base64 of "foo\nbar\n"
		got, err := fromBase64(data)
		require.NoError(t, err)
		require.Equal(t, "foo\nbar\n", got)
	})

	t.Run("invalid", func(t *testing.T) {
		data := "---"
		got, err := fromBase64(data)
		require.Error(t, err)
		require.Empty(t, got)
	})
}

func Test_toYaml(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		data := "foo"
		got, err := toYaml(&data)
		require.NoError(t, err)
		require.Equal(t, "foo\n", got)
	})

	t.Run("invalid", func(t *testing.T) {
		data := math.NaN()
		got, err := toYaml(&data)
		require.Error(t, err)
		require.Empty(t, got)
	})
}

func Test_writeCsv(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := writeCsv(nil, ',', "foo", "bar", "baz")
		require.NoError(t, err)
		require.Equal(t, "foo,bar,baz\n", got)
	})

	t.Run("write error", func(t *testing.T) {
		var bldr csvWriterBuilder = func(comma rune, buf *bytes.Buffer) csvWriter {
			return &dummyCsvWriter{
				write: func(record []string) error { return assert.AnError },
			}
		}
		got, err := writeCsv(bldr, ',', "foo", "bar", "baz")
		require.EqualError(t, err, assert.AnError.Error())
		require.Empty(t, got)
	})

	t.Run("flush error", func(t *testing.T) {
		var bldr csvWriterBuilder = func(comma rune, buf *bytes.Buffer) csvWriter {
			return &dummyCsvWriter{
				write: func(record []string) error { return nil },
				flush: func() {},
				error: func() error { return assert.AnError },
			}
		}
		got, err := writeCsv(bldr, ',', "foo", "bar", "baz")
		require.EqualError(t, err, assert.AnError.Error())
		require.Empty(t, got)
	})
}

type dummyCsvWriter struct {
	flush func()
	error func() error
	write func(record []string) error
}

func (d *dummyCsvWriter) Flush() {
	d.flush()
}

func (d *dummyCsvWriter) Error() error {
	return d.error()
}

func (d *dummyCsvWriter) Write(record []string) error {
	return d.write(record)
}
