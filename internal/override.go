package internal

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func ReposUploadReleaseAssetOverride(c *BaseCmd, filename string) error {
	if c.reqHeader == nil {
		c.reqHeader = http.Header{}
	}
	if !c.isValueSet("content-type") {
		mt, err := mimetype.DetectFile(filename)
		if err != nil {
			return err
		}
		c.reqHeader.Set("content-type", mt.String())
	}
	err := c.UseFileBody(filename)
	if err != nil {
		return err
	}
	c.curler = func(req *http.Request) (string, error) {
		req.Body = nil
		curl, err := defaultCurl(req)
		if err != nil {
			return "", err
		}
		*curl = append(*curl, "--data-binary", fmt.Sprintf("@%s", filename))
		return curl.String(), nil
	}

	//swap the url because this endpoint is actually at uploads.github.com
	c.APIBaseURL = strings.Replace(c.APIBaseURL, "api.github.com", "uploads.github.com", 1)
	// it must be an enterprise server if the url doesn't start with https://uploads.github.com
	// so we need to add /uploads to the beginning of the path
	if !strings.HasPrefix(c.APIBaseURL, "https://uploads.github.com") {
		c.APIBaseURL = path.Join(c.APIBaseURL, "uploads")
	}
	return nil
}
