// +build generator

package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

func main() {
	err := updateReadme("README.md")
	if err != nil {
		log.Fatal(err)
	}
}

var readmeRegexp = regexp.MustCompile(`(?s:<!--- START HELP OUTPUT --->.*<!--- END HELP OUTPUT --->)`)

func updateReadme(readmePath string) error {
	helpContent, err := getHelpOutput()
	if err != nil {
		return errors.Wrap(err, "failed getting help output")
	}

	oldReadmeContent, err := ioutil.ReadFile(readmePath)
	if err != nil {
		return errors.Wrapf(err, "failed reading file %q", readmePath)
	}
	newReadmeContent := readmeRegexp.ReplaceAll(oldReadmeContent, helpContent)

	return ioutil.WriteFile(readmePath, newReadmeContent, 0644)
}

func getHelpOutput() ([]byte, error) {
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
