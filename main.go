package main

import (
	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/services"
)

var version = "development"

type cli struct {
	services.CLI
	Version kong.VersionFlag
}

func main() {
	k := kong.Parse(&cli{}, kong.Vars{"version": version})
	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	k.FatalIfErrorf(k.Run(valueIsSetMap))
}
