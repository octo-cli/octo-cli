package main

import (
	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/generator"
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

type cli struct {
	Generate       generator.GenerateCmd       `cmd:"" help:"generate production code"`
	UpdateRoutes   generator.UpdateRoutesCmd   `cmd:"" help:"update routes.json with the latest"`
	UpdateTestdata generator.UpdateTestDataCmd `cmd:"" help:"updates routes.json and generated in generator/testdata"`
	UpdateReadme   generator.UpdateReadme      `cmd:"" help:"updates the help output section of README.md with whatever you pipe in here."`
	Build          build                       `cmd:"" help:"build bin/octo"`
	BuildLint      buildLint                   `cmd:"" help:"builds bin/golangci-lint"`
	Bootstrap      bootstrap                   `cmd:"" help:"bootstraps a dev environment"`
	Lint           lint                        `cmd:"" help:"run lint"`
}

func main() {
	k := kong.Parse(&cli{})
	k.FatalIfErrorf(k.Run())
}

type build struct{}

func (l *build) Run() error {
	return buildIfNeeded(filepath.Join(bindir, "octo"),
		".",
		"-s -w",
		true,
	)
}

type bootstrap struct {
	Force bool
}

func (l *bootstrap) Run() error {
	return buildIfNeeded(filepath.Join(bindir, "buildtool"),
		"./buildtool",
		"",
		l.Force)
}

type buildLint struct {
	Force     bool
	Version   string `default:"v1.12" help:"tag of golangci-lint to build"`
	BinTarget string `default:"./bin/golangci-lint" help:"where to put the golangci-lint binary"`
}

func (l *buildLint) Run() error {
	l.BinTarget = filepath.FromSlash(l.BinTarget)
	return buildGolangciLint(l.Version, l.BinTarget, l.Force)
}

type lint struct {
	Version string `default:"v1.12" help:"tag of golangci-lint to use"`
}

func (l *lint) Run() error {
	err := buildGolangciLint(l.Version, filepath.FromSlash("./bin/golangci-lint"), false)
	if err != nil {
		return errors.Wrap(err, "failed building golangci-lint")
	}
	cmd := exec.Command(filepath.FromSlash("./bin/golangci-lint"), "run", "--enable=golint")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
