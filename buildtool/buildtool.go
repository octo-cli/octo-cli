package main

import (
	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/generator"
)

type cli struct {
	Generate       generator.GenerateCmd       `cmd:"" help:"generate production code"`
	UpdateRoutes   generator.UpdateRoutesCmd   `cmd:"" help:"update routes.json with the latest"`
	UpdateTestdata generator.UpdateTestDataCmd `cmd:"" help:"updates routes.json and generated in generator/testdata"`
	UpdateReadme   generator.UpdateReadme      `cmd:"" help:"updates the help output section of README.md with whatever you pipe in here."`
}

func main() {
	k := kong.Parse(&cli{})
	k.FatalIfErrorf(k.Run())
}
