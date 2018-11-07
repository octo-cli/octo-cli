package main

import (
	"fmt"
	"github.com/golangci/golangci-lint/pkg/commands"
	"os"
)

func main() {
	e := commands.NewExecutor("", "", "")
	if err := e.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(1)
	}
}
