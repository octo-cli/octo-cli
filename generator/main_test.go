package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func Test_genCliRun_Run(t *testing.T) {
	wantIssuessvc, err := ioutil.ReadFile("testdata/exampleapp/services/issuessvc/issuessvc.go")
	require.Nil(t, err)
	err = os.MkdirAll("./tmp", 0755)
	require.Nil(t, err)
	tempDir, err := ioutil.TempDir("./tmp", "")
	defer func() {
		require.Nil(t, os.RemoveAll(tempDir))
	}()
	require.Nil(t, err)
	k := &genCliRun{
		RoutesPath: "testdata/routes.json",
		OutputPath: tempDir,
		ConfigFile: "testdata/exampleapp_config.hcl",
	}
	err = k.Run()
	assert.Nil(t, err)
	gotIssuessvc, err := ioutil.ReadFile(filepath.Join(tempDir, "services/issuessvc/issuessvc.go"))
	assert.Nil(t, err)
	assert.Equal(t, string(wantIssuessvc), string(gotIssuessvc))
}

