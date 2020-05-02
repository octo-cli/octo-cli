package main

import (
	"bufio"
	"bytes"
	"os"

	"github.com/pkg/errors"
)

func mustBeInRoot() error {
	wantLine := []byte("module github.com/octo-cli/octo-cli")
	modFile, err := os.Open("./go.mod")
	errMsg := "buildtool must be run from the root of octo-cli"
	if err != nil {
		return errors.New(errMsg)
	}
	defer func() {
		err = modFile.Close()
		if err != nil {
			panic(err)
		}
	}()
	bufReader := bufio.NewReader(modFile)
	line, _, err := bufReader.ReadLine()
	if err != nil {
		return errors.Wrap(err, "failed reading go.mod")
	}
	if !bytes.Equal(line, wantLine) {
		return errors.New(errMsg)
	}
	return nil
}
