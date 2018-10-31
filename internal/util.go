package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
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

//OutputResult writes the body of an http.Response to stdout
func OutputResult(resp *http.Response, rawJSON bool, stdout io.Writer) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if !rawJSON {
		body, err = prettyPrintJSON(body)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintln(stdout, string(body))
	return err
}
