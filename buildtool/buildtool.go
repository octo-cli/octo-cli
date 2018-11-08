package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/buildtool/generator"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	bindir string
)

func init() {
	bindir = filepath.Join(".", "bin")
}

type cli struct {
	Generate       generator.GenerateCmd       `cmd:"" help:"generate production code"`
	UpdateRoutes   UpdateRoutesCmd             `cmd:"" help:"update routes.json with the latest"`
	UpdateTestdata generator.UpdateTestDataCmd `cmd:"" help:"updates routes.json and generated in generator/testdata"`
	UpdateReadme   UpdateReadme                `cmd:"" help:"updates the help output section of README.md with whatever you pipe in here."`
	Build          build                       `cmd:"" help:"build bin/octo"`
	BuildLint      buildLint                   `cmd:"" help:"builds bin/golangci-lint"`
	Bootstrap      bootstrap                   `cmd:"" help:"bootstraps a dev environment"`
	Lint           lint                        `cmd:"" help:"run lint"`
	LatestVersion  latestVersionCmd            `cmd:"" help:"get the latest tagged version"`
	TagRelease     tagReleaseCmd               `cmd:"" help:"create a new version number"`
}

func main() {
	if err := mustBeInRoot(); err != nil {
		fmt.Println("buildtool must be run from the root of octo-cli")
		os.Exit(1)
	}
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

type latestVersionCmd struct {
	Nots bool `cmd:"" help:"strip any timestamp from the version"`
}

func (l *latestVersionCmd) Run() error {
	version, err := latestVersion(l.Nots)
	if err != nil {
		return err
	}
	fmt.Println(version.String())
	return nil
}

type tagReleaseCmd struct {
	Major      bool   `help:"new major version"`
	Minor      bool   `help:"new minor version"`
	Patch      bool   `help:"new patch version"`
	Prerelease string `help:"set prerelease"`
}

func (n *tagReleaseCmd) Run() error {
	tag, err := tagNewVersion(n.Major, n.Minor, n.Patch, n.Prerelease)
	if err != nil {
		return errors.Wrap(err, "failed tagging release")
	}
	fmt.Printf("tagging %q\n", tag)
	return nil
}

func (n *tagReleaseCmd) BeforeApply() error {
	var err error
	if !n.Major && !n.Minor && !n.Patch && n.Prerelease == "" {
		err = errors.New("You must set one of --major, --minor, --patch or --prerelease")
	}
	return err
}

type UpdateReadme struct {
	ReadmePath string `type:"existingfile" default:"README.md" help:"path to README.md"`
	Verify     bool   `help:"don't write anything.  Just verify README.md is current"`
}

func (k *UpdateReadme) Run() error {
	return updateReadme(k.ReadmePath, k.Verify)
}
