package jsonoutput

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// FormatJSONOutput formats json output with the template
func FormatJSONOutput(jsonBytes []byte, format, outputEach string) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, errors.New("invalid json")
	}
	if outputEach != "" {
		format = fmt.Sprintf("{{ range %s }}%s{{ end }}", outputEach, format)
	}
	var buf bytes.Buffer
	err = execTemplate(&buf, format, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func execTemplate(w io.Writer, pattern string, data interface{}) error {
	tmpl, err := template.New("").Funcs(funcMap).Parse(pattern)
	if err != nil {
		return errors.New("failed parsing format template")
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return errors.Wrap(err, "error executing template")
	}
	return nil
}

var funcMap = buildFuncMap()

var sprigFuncs = sprig.GenericFuncMap()

var directFromSprig = []string{
	// defaults
	"default", "empty", "coalesce", "ternary",

	// string
	"trim", "trimAll", "trimSuffix", "trimPrefix",
	"upper", "lower",
	"substr", "trunc", "contains", "replace",
	"cat",

	// string slice
	"sortAlpha",

	// type conversion
	"toString", "toStrings",

	// file path
	"base", "dir", "clean", "ext",

	// dicts
	"keys", "pick", "omit", "get", "hasKey",

	// lists
	"list",
}

func buildFuncMap() template.FuncMap {
	result := make(template.FuncMap, len(customFuncs)+len(directFromSprig))
	for i := 0; i < len(directFromSprig); i++ {
		result[directFromSprig[i]] = sprigFuncs[directFromSprig[i]]
	}
	for k, v := range customFuncs {
		result[k] = v
	}
	return result
}

var customFuncs = template.FuncMap{
	"newLine":      newLine,
	"fromBase64":   fromBase64,
	"pluck":        pluck,
	"join":         join,
	"toCsv":        csvFunc(','),
	"csvVals":      csvFunc(','),
	"toTsv":        csvFunc('\t'),
	"tsvVals":      csvFunc('\t'),
	"toYaml":       toYaml,
	"toJson":       sprigFuncs["mustToJson"],
	"toPrettyJson": sprigFuncs["mustToPrettyJson"],
	"toRawJson":    sprigFuncs["mustToRawJson"],
	"first":        sprigFuncs["mustFirst"],
	"last":         sprigFuncs["mustLast"],
	"uniq":         sprigFuncs["mustUniq"],
	"has":          sprigFuncs["mustHas"],
	"compact":      sprigFuncs["mustCompact"],
	"slice":        sprigFuncs["mustSlice"],
	"split":        sprigFuncs["splitList"],
}

func newLine(v interface{}) string {
	return sprigToString(v) + "\n"
}

func toYaml(v interface{}) (string, error) {
	b, err := yaml.Marshal(&v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func fromBase64(v string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func pluck(key string, vals ...interface{}) []interface{} {
	vals = expandSingleItemSlice(vals)
	res := []interface{}{}
	for _, v := range vals {
		dict, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if val, ok := dict[key]; ok {
			res = append(res, val)
		}
	}
	return res
}

func join(sep string, vals ...interface{}) string {
	return strings.Join(sprigToStrings(expandSingleItemSlice(vals)), sep)
}

func sprigToStrings(v interface{}) []string {
	fn := sprigFuncs["toStrings"].(func(interface{}) []string)
	return fn(v)
}

func sprigToString(v interface{}) string {
	fn := sprigFuncs["toString"].(func(interface{}) string)
	return fn(v)
}

func csvFunc(comma rune) func(vals ...interface{}) (string, error) {
	return func(vals ...interface{}) (string, error) {
		return writeCsv(nil, comma, vals...)
	}
}

type csvWriter interface {
	Flush()
	Error() error
	Write(record []string) error
}

type csvWriterBuilder func(comma rune, buf *bytes.Buffer) csvWriter

func newCsvWriter(comma rune, buf *bytes.Buffer) csvWriter {
	w := csv.NewWriter(buf)
	w.Comma = comma
	return w
}

func writeCsv(builder csvWriterBuilder, comma rune, vals ...interface{}) (string, error) {
	strVals := sprigToStrings(expandSingleItemSlice(vals))
	for i := range strVals {
		strVals[i] = strings.ReplaceAll(strVals[i], "\n", `\n`)
		strVals[i] = strings.ReplaceAll(strVals[i], "\t", `\t`)
		strVals[i] = strings.ReplaceAll(strVals[i], "\r", `\r`)
		strVals[i] = strings.ReplaceAll(strVals[i], `\\`, `\`)
	}
	if builder == nil {
		builder = newCsvWriter
	}
	var buf bytes.Buffer
	w := builder(comma, &buf)
	err := w.Write(strVals)
	if err != nil {
		return "", err
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// expandSingleItemSlice - if vals has a single item and that item is a slice or array, return v[0]
// otherwise just return vals
func expandSingleItemSlice(vals []interface{}) []interface{} {
	if len(vals) != 1 {
		return vals
	}
	val := reflect.ValueOf(vals[0])
	kind := val.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return vals
	}
	result := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		result[i] = val.Index(i).Interface()
	}
	return result
}
