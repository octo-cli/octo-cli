package generator

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

type (
	GenerateCmd struct {
		RoutesPath string `type:"existingfile" default:"routes.json"`
		OutputPath string `type:"existingdir" default:"./internal/generated"`
		Verify     bool   `help:"Verify a new run won't change anything"`
	}

	UpdateTestDataCmd struct{}
)

func (k *GenerateCmd) Run() error {
	if k.Verify {
		diffs, err := verify(k.RoutesPath, k.OutputPath)
		if err != nil {
			return errors.New("error verifying")
		}
		if len(diffs) > 0 {
			return fmt.Errorf("some files did not match: %v", diffs)
		}
	} else {
		Generate(k.RoutesPath, k.OutputPath)
	}
	return nil
}

func (k *UpdateTestDataCmd) Run() error {
	url := "https://octokit.github.io/routes/index.json"
	routesPath := "buildtool/generator/testdata/routes.json"
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

	Generate(routesPath, "buildtool/generator/testdata/generated")

	return errors.Wrap(err, "")
}
