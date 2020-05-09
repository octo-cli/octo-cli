package docgen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteDocs(t *testing.T) {
	fs := afero.NewMemMapFs()
	docsDir, err := afero.TempDir(fs, "", "")
	require.NoError(t, err)
	err = WriteDocs(filepath.FromSlash("../testdata/routes.json"), docsDir, fs)
	require.NoError(t, err)
	wantFiles, err := ioutil.ReadDir(filepath.FromSlash("../testdata/docs"))
	require.NoError(t, err)
	for _, wantFile := range wantFiles {
		wantData, err := ioutil.ReadFile(filepath.Join(filepath.FromSlash("../testdata/docs"), wantFile.Name()))
		require.NoError(t, err)
		gotData, err := afero.ReadFile(fs, filepath.Join(docsDir, wantFile.Name()))
		require.NoError(t, err)
		assert.Equalf(t, string(wantData), string(gotData), "file contents are not equal for %q", wantFile.Name())
	}
}
