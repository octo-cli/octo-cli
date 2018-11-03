package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_genCliRun_Run(t *testing.T) {
	t.Run("creates the right files", func(t *testing.T) {
		tempDir, cleanTempDir := createTempDir(t)
		defer cleanTempDir()
		k := &genCliRun{
			RoutesPath: "testdata/routes.json",
			OutputPath: tempDir,
		}
		err := k.Run()
		assert.Nil(t, err)
		genFiles := getDirectoryFileNames(t, tempDir)
		wantFiles := getDirectoryFileNames(t, "testdata/generated")
		assert.Equal(t, wantFiles, genFiles)
	})

	t.Run("file content matches", func(t *testing.T) {
		tempDir, cleanTempDir := createTempDir(t)
		defer cleanTempDir()
		k := &genCliRun{
			RoutesPath: "testdata/routes.json",
			OutputPath: tempDir,
		}
		err := k.Run()
		assert.Nil(t, err)
		wantFiles := getDirectoryFileNames(t, "testdata/generated")
		for _, wantFile := range wantFiles {
			wantData := readFile(t, "testdata/generated", wantFile)
			gotData := readFile(t, tempDir, wantFile)
			assert.Equalf(t, string(wantData), string(gotData), "file contents are not equal for %q", wantFile)
		}

	})
}

func readFile(t *testing.T, path ...string) []byte {
	t.Helper()
	data, err := ioutil.ReadFile(filepath.Join(path...))
	require.Nil(t, err)
	return data
}

func createTempDir(t *testing.T) (string, func()) {
	t.Helper()
	err := os.MkdirAll("./tmp", 0755)
	require.Nil(t, err)
	tempDir, err := ioutil.TempDir("./tmp", "")
	require.Nil(t, err)
	return tempDir, func() {
		t.Helper()
		require.Nil(t, os.RemoveAll(tempDir))
	}
}

func getDirectoryFileNames(t *testing.T, dir string) []string {
	t.Helper()
	var out []string
	files, err := ioutil.ReadDir(dir)
	require.Nil(t, err)
	for _, v := range files {
		out = append(out, v.Name())
	}
	sort.Strings(out)
	return out
}
