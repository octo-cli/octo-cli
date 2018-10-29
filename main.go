package main

import (
	"github.com/alecthomas/kong"
	"github.com/go-github-cli/go-github-cli/services"
)

func main() {
	cli := &services.CLI{}
	k := kong.Parse(cli)
	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	k.FatalIfErrorf(k.Run(valueIsSetMap))
}
