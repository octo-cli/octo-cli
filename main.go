package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/octo-cli/octo-cli/internal/generated"
)

var version = "development"

type cli struct {
	generated.CLI
	Version kong.VersionFlag
}

func main() {
	parser := kong.Must(&cli{},
		kong.Vars{"version": version},
		kong.Name("octo"),
	)

	for _, topLevelCmd := range parser.Model.Children {
		if topLevelCmd.Type != kong.CommandNode {
			continue
		}
		topLevelCmdHelps := generated.CmdHelps[topLevelCmd.Name]
		if topLevelCmdHelps == nil {
			topLevelCmdHelps = map[string]string{}
		}
		topLevelFlagHelps := generated.FlagHelps[topLevelCmd.Name]
		if topLevelFlagHelps == nil {
			topLevelFlagHelps = map[string]map[string]string{}
		}
		for _, cmd := range topLevelCmd.Children {
			if cmd.Type != kong.CommandNode {
				continue
			}
			if topLevelCmdHelps[cmd.Name] != "" {
				cmd.Help = topLevelCmdHelps[cmd.Name]
			}
			flagHelps := topLevelFlagHelps[cmd.Name]
			if flagHelps == nil {
				flagHelps = map[string]string{}
			}
			for _, flag := range cmd.Flags {
				if flagHelps[flag.Name] != "" {
					flag.Help = flagHelps[flag.Name]
				}
			}
		}
	}
	k, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	k.FatalIfErrorf(k.Run(valueIsSetMap))
}
