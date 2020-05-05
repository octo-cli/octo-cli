package generator

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type (
	GenerateCmd struct {
		RoutesPath string `type:"existingfile" default:"api.github.com.json"`
		OutputPath string `type:"existingdir" default:"./internal/generated"`
		DocsPath   string `type:"existingdir" default:"./docs"`
		Verify     bool   `help:"Verify a new run won't change anything"`
		fs         afero.Fs
	}
)

func (k *GenerateCmd) Run() error {
	if k.fs == nil {
		k.fs = afero.NewOsFs()
	}
	if k.Verify {
		diffs, err := verify(k.RoutesPath, k.OutputPath)
		if err != nil {
			return errors.New("error verifying")
		}
		if len(diffs) > 0 {
			return fmt.Errorf("some files did not match: %v", diffs)
		}
	} else {
		return Generate(k.RoutesPath, k.OutputPath, k.DocsPath, k.fs)
	}
	return nil
}
