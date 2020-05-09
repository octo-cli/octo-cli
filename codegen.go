// +build generator

package main

import (
	"flag"
	"log"

	"github.com/octo-cli/octo-cli/internal/generator/codegen"
	"github.com/spf13/afero"
)

func main() {
	var schema, outputPath string
	flag.StringVar(&schema, "schema", "", "")
	flag.StringVar(&outputPath, "out", "", "")
	flag.Parse()
	fs := afero.NewOsFs()
	err := codegen.Generate(schema, outputPath, fs)
	if err != nil {
		log.Fatal(err)
	}
}
