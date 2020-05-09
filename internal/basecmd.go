package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	"moul.io/http2curl/v2"
)

//UserAgent is the user agent for our https
var UserAgent = "octo-cli"

var (
	// Stdout is where to write output
	Stdout io.Writer = os.Stdout

	// Stderr is where to write error output
	Stderr io.Writer = os.Stderr

	// TransportWrapper is a wrapper for the http transport that we use for go-vcr tests
	TransportWrapper interface {
		SetTransport(t http.RoundTripper)
		http.RoundTripper
	}
)

// BaseCmd is included in all command structs in generated
type BaseCmd struct {
	isValueSetMap map[string]bool
	url           url.URL
	reqBody       map[string]interface{}
	acceptHeaders []string
	reqHeader     http.Header
	Token         string `env:"GITHUB_TOKEN" hidden:""`
	APIBaseURL    string `env:"GITHUB_API_BASE_URL" default:"https://api.github.com"`
	RawOutput     bool   `help:"don't format json output."`
	Format        string `help:"format json output with a go template"`
	Curl          bool   `help:"returns a corresponding curl request"`
}

func (c *BaseCmd) AfterApply() error {
	if c.Token == "" {
		return fmt.Errorf("missing environment variable: GITHUB_TOKEN")
	}
	return nil
}

//SetURLPath sets the path of url
func (c *BaseCmd) SetURLPath(path string) {
	c.url.Path = path
}

//SetIsValueSetMap sets isValueSetMap
func (c *BaseCmd) SetIsValueSetMap(isValueSetMap map[string]bool) {
	c.isValueSetMap = isValueSetMap
}

func (c *BaseCmd) isValueSet(valueName string) bool {
	return c.isValueSetMap[valueName]
}

func (c *BaseCmd) AddRequestHeader(headerName, value string) {
	if c.reqHeader == nil {
		c.reqHeader = http.Header{}
	}
	c.reqHeader.Add(headerName, value)
}

//UpdateBody adds a flag's value a request body
func (c *BaseCmd) UpdateBody(flagName string, value interface{}) {
	if c.reqBody == nil {
		c.reqBody = map[string]interface{}{}
	}
	if c.isValueSet(flagName) {
		key := strings.Split(flagName, ".")
		setBodyValue(c.reqBody, key, value)
	}
}

func setBodyValue(body map[string]interface{}, key []string, value interface{}) {
	if len(key) == 1 {
		body[key[0]] = value
		return
	}
	var nextMap map[string]interface{}
	switch currentVal := body[key[0]].(type) {
	case map[string]interface{}:
		nextMap = currentVal
	default:
		nextMap = map[string]interface{}{}
		body[key[0]] = nextMap
	}
	setBodyValue(nextMap, key[1:], value)
}

//UpdateURLPath sets a param in the url path
func (c *BaseCmd) UpdateURLPath(valName string, value interface{}) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	case int64:
		strVal = strconv.FormatInt(v, 10)
	case nil:
		strVal = ""
	}
	c.url.Path = strings.Replace(c.url.Path, "{"+valName+"}", strVal, 1)
}

//UpdatePreview adds a preview header to a request
func (c *BaseCmd) UpdatePreview(previewName string, value bool) {
	if value {
		accept := fmt.Sprintf("application/vnd.github.%s-preview+json", previewName)
		c.acceptHeaders = append(c.acceptHeaders, accept)
	}
}

//UpdateURLQuery sets a param value in a url query
func (c *BaseCmd) UpdateURLQuery(paramName string, value interface{}) {
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

func (c *BaseCmd) newRequest(method string) (*http.Request, error) {
	c.url.Path = strings.Replace(c.url.Path, "//", "/", -1)
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
	req.Header.Set("User-Agent", UserAgent)
	reqHeader := c.reqHeader.Clone()
	for k, v := range reqHeader {
		req.Header[k] = v
	}
	return req, nil
}

func (c *BaseCmd) httpClient() *http.Client {
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token}))
	if TransportWrapper != nil {
		TransportWrapper.SetTransport(tc.Transport)
		tc.Transport = TransportWrapper
	}
	return tc
}

//DoRequest performs a request
func (c *BaseCmd) DoRequest(method string) error {
	req, err := c.newRequest(method)
	if err != nil {
		return err
	}

	if c.Curl {
		var curl *http2curl.CurlCommand
		curl, err = http2curl.GetCurlCommand(req)
		if err != nil {
			return err
		}
		for i, curler := range *curl {
			if i == 0 || (*curl)[i-1] != "-d" {
				continue
			}
			(*curl)[i] = regexp.MustCompile(`\s*'\s*$`).ReplaceAllString(curler, "'")
		}
		*curl = append((*curl)[:len(*curl)-1],
			"-H",
			`"Authorization: token $GITHUB_TOKEN"`,
			(*curl)[len(*curl)-1],
		)
		fmt.Fprintln(Stdout, curl.String())
		return nil
	}

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}

	return OutputResult(resp, c.RawOutput, c.Format, Stdout)
}
