package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func buildIfNeeded(bin, src, ldflags string, force bool) error {
	if err := mustBeInRoot(); err != nil {
		return err
	}
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
	ok, err := fileExists("./buildtool/util.go")
	if err != nil {
		return errors.Wrap(err, "failed to determine if we are in the right directory")
	}
	if !ok {
		return errors.New("this must be run from the root of octo-cli")
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
	if err := mustBeInRoot(); err != nil {
		return err
	}
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
