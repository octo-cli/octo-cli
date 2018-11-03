package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/octo-cli/octo-cli/internal"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Stdout is where to write output
var Stdout io.Writer = os.Stdout

// Stderr is where to write error output
var Stderr io.Writer = os.Stderr

// TransportWrapper is a wrapper for the http transport that we use for go-vcr tests
var TransportWrapper interface {
	SetTransport(t http.RoundTripper)
	http.RoundTripper
}

type baseCmd struct {
	isValueSetMap map[string]bool
	url           url.URL
	reqBody       *map[string]interface{}
	acceptHeaders []string
	Token         string `env:"GITHUB_TOKEN" required:""`
	APIBaseURL    string `env:"GITHUB_API_BASE_URL" default:"https://api.github.com"`
	RawOutput     bool   `help:"don't format json output."`
	Format        string `help:"format output with a go template"`
}

func (c *baseCmd) isValueSet(valueName string) bool {
	return c.isValueSetMap[valueName]
}

func (c *baseCmd) updateBody(flagName string, value interface{}) {
	if c.reqBody == nil {
		c.reqBody = &map[string]interface{}{}
	}
	if c.isValueSet(flagName) {
		b := *c.reqBody
		b[flagName] = value
		c.reqBody = &b
	}
}

func (c *baseCmd) updateURLPath(valName string, value interface{}) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	case int64:
		strVal = strconv.FormatInt(v, 10)
	case nil:
		strVal = ""
	}
	c.url.Path = strings.Replace(c.url.Path, ":"+valName, strVal, 1)
}

func (c *baseCmd) updatePreview(previewName string, value bool) {
	if value {
		accept := fmt.Sprintf("application/vnd.github.%s-preview+json", previewName)
		c.acceptHeaders = append(c.acceptHeaders, accept)
	}
}

func (c *baseCmd) updateURLQuery(paramName string, value interface{}) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	case int64:
		strVal = strconv.FormatInt(v, 10)
	}
	if c.isValueSet(paramName) {
		query := c.url.Query()
		query.Add(paramName, strVal)
		c.url.RawQuery = query.Encode()
	}
}

func (c *baseCmd) newRequest(method string) (*http.Request, error) {
	u := strings.Join([]string{
		strings.TrimSuffix(c.APIBaseURL, "/"),
		strings.TrimPrefix(c.url.String(), "/"),
	}, "/")
	var buf io.ReadWriter
	if c.reqBody != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(c.reqBody)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}
	if c.reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	acceptHeaders := []string{"application/vnd.github.v3+json"}
	acceptHeaders = append(acceptHeaders, c.acceptHeaders...)
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	return req, nil
}

func (c *baseCmd) httpClient() *http.Client {
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token}))
	if TransportWrapper != nil {
		TransportWrapper.SetTransport(tc.Transport)
		tc.Transport = TransportWrapper
	}
	return tc
}

func (c *baseCmd) doRequest(method string) error {
	req, err := c.newRequest(method)
	if err != nil {
		return err
	}

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}

	return internal.OutputResult(resp, c.RawOutput, c.Format, Stdout)
}
