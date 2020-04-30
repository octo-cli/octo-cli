package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		err = modFile.Close()
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
	tdRoutesPath := "buildtool/generator/testdata/routes.json"
	err := copyFile("routes.json", tdRoutesPath)
	if err != nil {
		return errors.Wrap(err, "failed copying routes.json")
	}
	generator.Generate(tdRoutesPath, "buildtool/generator/testdata/generated", nil)
	return nil
}

func copyFile(srcRoutesPath string, tdRoutesPath string) error {
	routes, err := ioutil.ReadFile(srcRoutesPath)
	if err != nil {
		return errors.Wrapf(err, "failed reading %q", srcRoutesPath)
	}
	err = ioutil.WriteFile(tdRoutesPath, routes, 0644)
	return errors.Wrapf(err, "failed writing %q", tdRoutesPath)
}
