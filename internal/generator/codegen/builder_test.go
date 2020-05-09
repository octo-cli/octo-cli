package codegen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	fs := afero.NewMemMapFs()
	genDir, err := afero.TempDir(fs, "", "")
	require.NoError(t, err)
	err = Generate(filepath.FromSlash("../testdata/routes.json"), genDir, fs)
	require.NoError(t, err)

	wantFiles, err := ioutil.ReadDir(filepath.FromSlash("../testdata/generated"))
	require.NoError(t, err)
	for _, wantFile := range wantFiles {
		wantData, err := ioutil.ReadFile(filepath.Join(filepath.FromSlash("../testdata/generated"), wantFile.Name()))
		require.NoError(t, err)
		gotData, err := afero.ReadFile(fs, filepath.Join(genDir, wantFile.Name()))
		require.NoError(t, err)
		assert.Equalf(t, string(wantData), string(gotData), "file contents are not equal for %q", wantFile.Name())
	}
}
