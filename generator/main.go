package generator

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const routesTimestampFormat = "20060102T150405Z0700"

type (
	UpdateRoutesCmd struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	GenerateCmd struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		OutputPath string `type:"existingdir" default:"./internal/generated"`
		Verify     bool   `help:"Verify a new run won't change anything"`
	}

	UpdateReadme struct {
		ReadmePath string `type:"existingfile" default:"README.md" help:"path to README.md"`
		Verify     bool   `help:"don't write anything.  Just verify README.md is current"`
	}

	UpdateTestDataCmd struct{}
)

var readmeRegexp = regexp.MustCompile(`(?s:<!--- START HELP OUTPUT --->.*<!--- END HELP OUTPUT --->)`)

func getHelpOutput() ([]byte, error) {
	if _, err := os.Stat("routes.json"); err != nil {
		return nil, errors.New("must be run from octo-cli root")
	}
	cmd := exec.Command("go", "run", ".", "--help")
	helpOut, err := cmd.Output()
	switch e := err.(type) {
	case nil:
	case *exec.ExitError:
		if e.Error() != "exit status 1" {
			return nil, err
		}
	default:
		return nil, err
	}
	output := []byte("<!--- START HELP OUTPUT --->\n```\n")
	output = append(output, helpOut...)
	output = append(output, []byte("\n```\n<!--- END HELP OUTPUT --->")...)
	return output, nil
}

func (k *UpdateReadme) Run() error {
	helpContent, err := getHelpOutput()
	if err != nil {
		return errors.Wrap(err, "failed getting help output")
	}

	oldReadmeContent, err := ioutil.ReadFile(k.ReadmePath)
	if err != nil {
		return errors.Wrapf(err, "failed reading file %q", k.ReadmePath)
	}
	newReadmeContent := readmeRegexp.ReplaceAll(oldReadmeContent, helpContent)

	if k.Verify {
		if !bytes.Equal(newReadmeContent, oldReadmeContent) {
			err = errors.Errorf("%q is not current", k.ReadmePath)
		}
	} else {
		err = ioutil.WriteFile(k.ReadmePath, newReadmeContent, 0644)
		err = errors.Wrapf(err, "failed writing file %q", k.ReadmePath)
	}
	return err
}

func (k *GenerateCmd) Run() error {
	if k.Verify {
		diffs, err := verify(k.RoutesPath, k.OutputPath)
		if err != nil {
			return errors.New("error verifying")
		}
		if len(diffs) > 0 {
			return fmt.Errorf("some files did not match: %v", diffs)
		}
	} else {
		Generate(k.RoutesPath, k.OutputPath)
	}
	return nil
}

func updateRoutes(routesURL, routesPath string) error {
	resp, err := http.Get(routesURL)
	if err != nil {
		return errors.Wrapf(err, "failed getting %q", routesURL)
	}

	err = writeRoutesJSON(routesPath, resp)
	if err != nil {
		return errors.Wrap(err, "failed writing routesJSON")
	}

	err = updateLastModified(routesPath, resp)
	return errors.Wrap(err, "failed updateing last modified")
}

func writeRoutesJSON(routesPath string, resp *http.Response) error {
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", routesPath)
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", routesPath)
	}
	err = resp.Body.Close()
	if err != nil {
		return errors.Wrap(err, "failed closing response body")
	}
	return errors.Wrapf(err, "failed closing file %q", outFile.Close())
}

func updateLastModified(routesPath string, resp *http.Response) error {
	extension := filepath.Ext(routesPath)
	if extension == "" {
		return errors.Errorf("routesPath must have a file extension (preferable .json)")
	}
	lmPath := strings.TrimSuffix(routesPath, extension) + "-last-modified.txt"
	lmHeader := resp.Header.Get("last-modified")
	lmTime, err := time.Parse(time.RFC1123, lmHeader)
	if err != nil {
		return errors.Wrapf(err, "couldn't parse last-modified time %q", lmHeader)
	}
	lmFile, err := os.Create(lmPath)
	if err != nil {
		return errors.Wrapf(err, "failed creating file %q", lmPath)
	}
	_, err = lmFile.WriteString(lmTime.Format(routesTimestampFormat))
	if err != nil {
		return errors.Wrapf(err, "failed writing to file %q", lmPath)
	}
	return errors.Wrapf(lmFile.Close(), "failed closing file %q", lmPath)
}

func (k *UpdateRoutesCmd) Run() error {
	return updateRoutes(k.RoutesURL, k.RoutesPath)
}

func (k *UpdateTestDataCmd) Run() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "generator/testdata/routes.json"
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "")
	}
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}

	Generate(routesPath, "generator/testdata/generated")

	return errors.Wrap(err, "")
}
