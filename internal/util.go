package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"text/template"

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

func formatOutput(jsonBytes []byte, format string) ([]byte, error) {
	var it interface{}
	err := json.Unmarshal(jsonBytes, &it)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling json")
	}
	tmpl, err := template.New("").Parse(format)
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
	contentType, _, err := mime.ParseMediaType(resp.Header.Get("content-type"))
	_ = err //nolint:errcheck //just treat it like raw text if we can't get a media type

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	switch {
	case format != "":
		body, err = formatOutput(body, format)
		if err != nil {
			return err
		}
	case contentType == "application/json" && !rawOutput:
		body, err = prettyPrintJSON(body)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintln(stdout, string(body))
	return err
}
