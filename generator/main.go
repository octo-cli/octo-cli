package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

type (
	genCli struct {
		Run            genCliRun            `cmd:"" help:"generate production code"`
		UpdateRoutes   genCliUpdateRoutes   `cmd:"" help:"update routes.json with the latest"`
		UpdateTestdata genCliUpdateTestdata `cmd:"" help:"updates routes.json and services in generator/testdata"`
	}

	genCliUpdateRoutes struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		RoutesURL  string `default:"https://octokit.github.io/routes/index.json"`
	}

	genCliRun struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		OutputPath string `type:"existingdir" default:"./services"`
	}

	genCliUpdateTestdata struct{}
)

func (k *genCliRun) Run() error {
	Generate(k.RoutesPath, k.OutputPath)
	return nil
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

	Generate(routesPath, "generator/testdata/services")

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
