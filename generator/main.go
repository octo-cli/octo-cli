package main

import (
	"fmt"
	"github.com/WillAbides/go-github-cli/generator/internal"
	"github.com/WillAbides/go-github-cli/generator/internal/generator"
	"github.com/WillAbides/go-github-cli/generator/internal/packagewriter"
	"io"
	"net/http"
	"os"

	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
)

type (
	genCli struct {
		Run            genCliRun            `cmd:"" help:"generate production code"`
		UpdateRoutes   genCliUpdateRoutes   `cmd:"" help:"update routes.json with the latest"`
		UpdateTestdata genCliUpdateTestdata `cmd:"" help:"updates routes.json and exampleapp in generator/testdata"`
	}

	genCliUpdateRoutes struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	genCliRun struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		ConfigFile string `type:"existingfile" default:"config.hcl"`
		OutputPath string `type:"existingdir" default:"."`
	}

	genCliUpdateTestdata struct{}
)

func (k *genCliRun) Run() error {
	svcs, err := generator.BuildSvcs(k.RoutesPath, k.ConfigFile)
	if err != nil {
		return errors.Wrap(err, "")
	}
	for _, svc := range svcs {
		err = packagewriter.WritePackageFiles(k.OutputPath, svc)
	}
	var toPkgers []interface{ ToPkg() (*internal.Pkg, error) }
	for _, v := range svcs {
		toPkgers = append(toPkgers, v)
		if err != nil {
			return errors.Wrap(err, "failed writing package files")
		}
	}
	return errors.Wrap(err, "")
}

func (k *genCliUpdateRoutes) Run() error {
	resp, err := http.Get(k.RoutesURL)
	if err != nil {
		return errors.Wrap(err, "")
	}
	outFile, err := os.Create(k.RoutesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	return errors.Wrap(err, "")
}

func (k *genCliUpdateTestdata) Run() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "generator/testdata/routes.json"
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "")
	}
	outFile, err := os.Create(routesPath)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	svcs, err := generator.BuildSvcs(routesPath, "generator/testdata/exampleapp_config.hcl")
	if err != nil {
		return errors.Wrap(err, "")
	}
	for _, svc := range svcs {
		err = packagewriter.WritePackageFiles("generator/testdata/exampleapp", svc)
		if err != nil {
			return errors.Wrap(err, "failed writing package files")
		}
	}
	return errors.Wrap(err, "")
}

func main() {
	k := kong.Parse(&genCli{})
	err := k.Run()
	if err != nil {
		fmt.Printf("%+v", err)
		fmt.Println(err)
	}
	k.FatalIfErrorf(err)
}
