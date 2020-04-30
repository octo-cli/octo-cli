package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/buildtool/generator"
	"github.com/pkg/errors"
)

var (
	bindir string
)

func init() {
	bindir = filepath.Join(".", "bin")
}

type cli struct {
	Build               build                  `cmd:"" help:"build bin/octo"`
	Generate            generator.GenerateCmd  `cmd:"" help:"generate production code"`
	LatestTaggedRelease latestTaggedReleaseCmd `cmd:"" help:"get the latest tagged version"`
	TagRelease          tagReleaseCmd          `cmd:"" help:"creates a git tag for a new release of octo-cli"`
	UpdateReadme        updateReadmeCmd        `cmd:"" help:"updates the help output section of README.md"`
	UpdateRoutes        updateRoutesCmd        `cmd:"" help:"update routes.json with the latest"`
	UpdateTestdata      updateTestDataCmd      `cmd:"" help:"updates routes.json and generated in internal/generator/testdata"`
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

type latestTaggedReleaseCmd struct {
	Nots bool `cmd:"" help:"strip any timestamp from the version"`
}

func (l *latestTaggedReleaseCmd) Run() error {
	version, err := latestTaggedRelease(l.Nots)
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

func (n *tagReleaseCmd) AfterApply() error {
	var err error
	if !n.Major && !n.Minor && !n.Patch && n.Prerelease == "" {
		err = errors.New("You must set one of --major, --minor, --patch or --prerelease")
	}
	return err
}

type updateReadmeCmd struct {
	ReadmePath string `type:"existingfile" default:"README.md" help:"path to README.md"`
	Verify     bool   `help:"don't write anything.  Just verify README.md is current"`
}

func (k *updateReadmeCmd) Run() error {
	return updateReadme(k.ReadmePath, k.Verify)
}

type updateRoutesCmd struct {
	RoutesPath string `type:"existingfile" default:"routes.json"`
	RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
}

func (k *updateRoutesCmd) Run() error {
	return updateRoutes(k.RoutesURL, k.RoutesPath)
}

type updateTestDataCmd struct{}

func (k *updateTestDataCmd) Run() error {
	return updateTestData()
}
