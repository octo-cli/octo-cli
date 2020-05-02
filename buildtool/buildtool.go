package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/buildtool/generator"
)

type cli struct {
	Generate     generator.GenerateCmd `cmd:"" help:"generate production code"`
	UpdateReadme updateReadmeCmd       `cmd:"" help:"updates the help output section of README.md"`
}

func main() {
	if err := mustBeInRoot(); err != nil {
		fmt.Println("buildtool must be run from the root of octo-cli")
		os.Exit(1)
	}
	k := kong.Parse(&cli{})
	k.FatalIfErrorf(k.Run())
}

type updateReadmeCmd struct {
	ReadmePath string `type:"existingfile" default:"README.md" help:"path to README.md"`
	Verify     bool   `help:"don't write anything.  Just verify README.md is current"`
}

func (k *updateReadmeCmd) Run() error {
	return updateReadme(k.ReadmePath, k.Verify)
}
