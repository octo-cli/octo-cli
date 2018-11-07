package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	bindir    string
	scriptdir string
)

func init() {
	bindir = filepath.Join(".", "bin")
	scriptdir = filepath.Join(".", "script")
}

func main() {
	k := kong.Parse(&cli{})
	k.FatalIfErrorf(k.Run())
}

type cli struct {
	BuildLint buildLint `kong:"cmd"`
	Bootstrap bootstrap `kong:"cmd"`
	Build     build     `kong:"cmd"`
}

type buildLint struct {
	Force bool
}

func (l *buildLint) Run() error {
	return buildIfNeeded(filepath.Join(bindir, "golangci-lint"),
		filepath.Join(scriptdir, "util", "golangci-lint", "main.go"),
		"",
		l.Force)
}

type bootstrap struct {
	Force bool
}

func (l *bootstrap) Run() error {
	return buildIfNeeded(filepath.Join(bindir, "util"),
		filepath.Join(scriptdir, "util", "util.go"),
		"",
		l.Force)
}

type build struct{}

func (l *build) Run() error {
	return buildIfNeeded(filepath.Join(bindir, "octo"),
		".",
		"-s -w",
		true,
	)
}

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
	ok, err := fileExists("./script/util/util.go")
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
