package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func prettyPrintJSON(jsonBytes []byte) ([]byte, error) {
	var it interface{}
	err := json.Unmarshal(jsonBytes, &it)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling json")
	}
	out, err := json.MarshalIndent(it, "", "  ")
	return out, errors.Wrap(err, "failed marshalling json")
}

var templateFuncs = template.FuncMap{
	"json": func(v interface{}) (string, error) {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(v)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(buf.String()), nil
	},
	"yaml": func(v interface{}) (string, error) {
		r, err := yaml.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(r), nil
	},
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
}

func formatOutput(jsonBytes []byte, format, outputEach string) ([]byte, error) {
	var it interface{}
	err := json.Unmarshal(jsonBytes, &it)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling json")
	}
	format = strings.ReplaceAll(format, `\n`, "\n")
	format = strings.ReplaceAll(format, `\t`, "\t")
	if outputEach != "" {
		format = fmt.Sprintf("{{ range %s }}%s{{ end}}", outputEach, format)
	}
	tmpl, err := template.New("").Funcs(templateFuncs).Parse(format)
	if err != nil {
		return nil, errors.Wrap(err, "failed parsing format template")
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, it)
	if err != nil {
		return nil, errors.Wrap(err, "failed executing template")
	}
	return buf.Bytes(), err
}
