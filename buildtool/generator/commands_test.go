package generator

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testdataFs = afero.NewOsFs()

func Test_genCliRun_Run(t *testing.T) {
	t.Run("file content matches", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		genDir, err := afero.TempDir(fs, "", "")
		require.NoError(t, err)
		docsDir, err := afero.TempDir(fs, "", "")
		require.NoError(t, err)
		k := &GenerateCmd{
			RoutesPath: "testdata/routes.json",
			OutputPath: genDir,
			DocsPath:   docsDir,
			fs:         fs,
		}
		err = k.Run()
		assert.NoError(t, err)
		wantFiles := getDirectoryFileNames(t, "testdata/generated", testdataFs)
		for _, wantFile := range wantFiles {
			wantData := readFile(t, testdataFs, "testdata/generated", wantFile)
			gotData := readFile(t, fs, genDir, wantFile)
			assert.Equalf(t, string(wantData), string(gotData), "file contents are not equal for %q", wantFile)
		}
		wantUnsup, err := ioutil.ReadFile("testdata/docs/unsupported.md")
		require.NoError(t, err)
		gotUnsup := readFile(t, fs, docsDir, "unsupported.md")
		require.Equal(t, wantUnsup, gotUnsup)
		wantOperations, err := ioutil.ReadFile("testdata/docs/operations.md")
		require.NoError(t, err)
		gotOps := readFile(t, fs, docsDir, "operations.md")
		require.Equal(t, wantOperations, gotOps)
	})
}

func readFile(t *testing.T, fs afero.Fs, path ...string) []byte {
	t.Helper()
	data, err := afero.ReadFile(fs, filepath.Join(path...))
	require.Nil(t, err)
	return data
}

func getDirectoryFileNames(t *testing.T, dir string, fs afero.Fs) []string {
	t.Helper()
	var out []string
	files, err := afero.ReadDir(fs, dir)
	require.Nil(t, err)
	for _, v := range files {
		out = append(out, v.Name())
	}
	sort.Strings(out)
	return out
}
