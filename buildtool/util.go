package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/masterminds/semver"
	"github.com/octo-cli/octo-cli/buildtool/generator"
	"github.com/pkg/errors"
)

func buildIfNeeded(bin, src, ldflags string, force bool) error {
	ok, err := fileExists(bin)
	if err != nil {
		return errors.Wrapf(err, "failed finding %q", bin)
	}
	if ok && !force {
		return nil
	}
	args := []string{"build", "-o", bin, src}
	if ldflags != "" {
		args = []string{"build", fmt.Sprintf("-ldflags=%s", ldflags), "-o", bin, src}
	}
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	}
	return errors.Wrapf(err, "failed building %q", bin)
}

func mustBeInRoot() error {
	wantLine := []byte("module github.com/octo-cli/octo-cli")
	modFile, err := os.Open("./go.mod")
	errMsg := "buildtool must be run from the root of octo-cli"
	if err != nil {
		return errors.New(errMsg)
	}
	defer func() {
		err := modFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	bufReader := bufio.NewReader(modFile)
	line, _, err := bufReader.ReadLine()
	if err != nil {
		return errors.Wrap(err, "failed reading go.mod")
	}
	if !bytes.Equal(line, wantLine) {
		return errors.New(errMsg)
	}
	return nil
}

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func buildGolangciLint(tag string, outputPath string, force bool) error {
	ok, err := fileExists(outputPath)
	if err != nil {
		return errors.Wrapf(err, "failed finding %q", outputPath)
	}
	if ok && !force {
		return nil
	}
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			panic(err)
		}
	}()
	outputPath, err = filepath.Abs(outputPath)
	if err != nil {
		return errors.Wrapf(err, "failed getting absolute path of %q", outputPath)
	}
	cloneCmd := exec.Command("git", "clone", "--branch", tag, "https://github.com/golangci/golangci-lint", dir)
	err = cloneCmd.Run()
	if err != nil {
		return errors.Wrap(err, "failed cloning golangci-lint")
	}
	buildpath := strings.Join([]string{".", "cmd", "golangci-lint"}, string(os.PathSeparator))
	buildCmd := exec.Command("go", "build", "-ldflags", "-s -w", "-o", outputPath, buildpath)
	buildCmd.Dir = dir
	err = buildCmd.Run()
	return errors.Wrap(err, "failed building golangci-linit")
}

func latestTaggedRelease(stripPre bool) (*semver.Version, error) {
	tagBytes, err := exec.Command("git", "describe", "--tags", "--match", "v*[0-9].*[0-9].*[0-9]*", "--abbrev=0").Output()
	if err != nil {
		return nil, errors.Wrap(err, "could not find tag")
	}
	tag := string(tagBytes)
	tag = strings.TrimSpace(tag)
	tag = strings.TrimPrefix(tag, "v")
	version, err := semver.NewVersion(tag)
	if err != nil {
		return nil, errors.Wrap(err, "failed parsing version")
	}
	if stripPre {
		newVer, err := version.SetPrerelease("")
		if err != nil {
			return nil, err
		}
		version = &newVer
	}
	return version, nil
}

func tagNewVersion(major, minor, patch bool, prerelease string) (string, error) {
	tag, err := newVersionTag(major, minor, patch, prerelease)
	if err != nil {
		return "", errors.Wrap(err, "failed creating new tag name")
	}
	return tag, exec.Command("git", "tag", tag).Run()
}

func newVersionTag(major, minor, patch bool, prerelease string) (string, error) {
	nextVersion, err := latestTaggedRelease(true)
	if err != nil {
		return "", errors.Wrap(err, "failed getting latest version")
	}
	version := *nextVersion
	if patch {
		version = version.IncPatch()
	}
	if minor {
		version = version.IncMinor()
	}
	if major {
		version = version.IncMajor()
	}
	version, err = version.SetPrerelease(prerelease)
	if err != nil {
		return "", errors.Wrap(err, "failed setting prerelease")
	}
	return "v" + version.String(), nil
}

func updateTestData() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "buildtool/generator/testdata/routes.json"
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

	generator.Generate(routesPath, "buildtool/generator/testdata/generated", nil)

	return errors.Wrap(err, "")
}

func cibuild() error {
	var failures []string
	for _, script := range [][]string{
		{"script/test", "-race"},
		{"script/lint"},
		{"script/generate", "--verify"},
		{"script/update-readme", "--verify"},
	} {
		fmt.Printf("\nrunning %s\n\n", strings.Join(script, " "))
		cmd := exec.Command(script[0])
		if len(script) > 1 {
			cmd = exec.Command(script[0], script[1:]...)
		}
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			failures = append(failures, strings.Join(script, " "))
		}
	}
	if len(failures) > 0 {
		return errors.New("cibuild failed")
	}
	return nil
}
