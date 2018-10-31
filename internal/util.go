package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
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

func formatOutput(jsonBytes []byte, format string) ([]byte, error) {
	var it interface{}
	err := json.Unmarshal(jsonBytes, &it)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling json")
	}
	tmpl, err := template.New("").Parse(string(format))
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

//OutputResult writes the body of an http.Response to stdout
func OutputResult(resp *http.Response, rawOutput bool, format string, stdout io.Writer) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if format != "" {
		body, err = formatOutput(body, format)
		if err != nil {
			return err
		}
	} else {
		if !rawOutput {
			body, err = prettyPrintJSON(body)
			if err != nil {
				return err
			}
		}
	}
	_, err = fmt.Fprintln(stdout, string(body))
	return err
}
