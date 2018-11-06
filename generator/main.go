package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const routesTimestampFormat = "20060102T150405Z0700"

type (
	genCli struct {
		Run              genCliRun              `cmd:"" help:"generate production code"`
		UpdateRoutes     genCliUpdateRoutes     `cmd:"" help:"update routes.json with the latest"`
		UpdateTestdata   genCliUpdateTestdata   `cmd:"" help:"updates routes.json and generated in generator/testdata"`
		UpdateReadmeHelp genCliUpdateReadmeHelp `cmd:"" help:"updates the help output section of README.md with whatever you pipe in here."`
	}

	genCliUpdateRoutes struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	genCliRun struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		OutputPath string `type:"existingdir" default:"./internal/generated"`
		Verify     bool   `help:"Verify a new run won't change anything"`
	}

	genCliUpdateReadmeHelp struct {
		ReadmePath string `type:"existingfile" default:"README.md" help:"path to README.md"`
	}

	genCliUpdateTestdata struct{}
)

var readmeRegexp = regexp.MustCompile(`(?s:<!--- START HELP OUTPUT --->.*<!--- END HELP OUTPUT --->)`)

func (k *genCliUpdateReadmeHelp) Run() error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return errors.Errorf("generator update-readme-help is intended to work with pipes.\nUsage: echo \"whatever\" | generator update-readme-help")
	}

	reader := bufio.NewReader(os.Stdin)
	var helpContent []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		helpContent = append(helpContent, input)
	}
	newHelp := "<!--- START HELP OUTPUT --->\n```\n" +
		string(helpContent) +
		"\n```\n<!--- END HELP OUTPUT --->"

	readmeBytes, err := ioutil.ReadFile(k.ReadmePath)
	if err != nil {
		return errors.Wrapf(err, "failed reading file %q", k.ReadmePath)
	}
	newReadmeContent := readmeRegexp.ReplaceAll(readmeBytes, []byte(newHelp))
	err = ioutil.WriteFile(k.ReadmePath, newReadmeContent, 0644)
	return errors.Wrapf(err, "failed writing file %q", k.ReadmePath)
}

func (k *genCliRun) Run() error {
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

func (k *genCliUpdateRoutes) Run() error {
	return updateRoutes(k.RoutesURL, k.RoutesPath)
}

func (k *genCliUpdateTestdata) Run() error {
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

func main() {
	k := kong.Parse(&genCli{})
	err := k.Run()
	k.FatalIfErrorf(err)
}
