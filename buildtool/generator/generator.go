package generator

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/octo-cli/octo-cli/internal/generator/codegen"
	"github.com/octo-cli/octo-cli/internal/generator/docgen"
	"github.com/spf13/afero"
)

func dirFileMap(path string, fs afero.Fs) (map[string][]byte, error) {
	output := map[string][]byte{}
	fileInfos, err := afero.ReadDir(fs, path)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fileInfo.Sys()
		if fileInfo.IsDir() {
			continue
		}
		fileBytes, err := afero.ReadFile(fs, filepath.Join(path, fileInfo.Name()))
		if err != nil {
			return nil, err
		}
		output[fileInfo.Name()] = fileBytes
	}
	return output, nil
}

func verify(routesPath, outputPath string) ([]string, error) {
	realFs := afero.NewOsFs()
	tmpFs := afero.NewMemMapFs()
	tempDir, err := afero.TempDir(tmpFs, "", "")
	if err != nil {
		return nil, err
	}
	tmpDocsDir, err := afero.TempDir(tmpFs, "", "")
	if err != nil {
		return nil, err
	}
	err = Generate(routesPath, tempDir, tmpDocsDir, tmpFs)
	if err != nil {
		return nil, err
	}

	wantFiles, err := dirFileMap(tempDir, tmpFs)
	if err != nil {
		return nil, err
	}
	gotFiles, err := dirFileMap(outputPath, realFs)
	if err != nil {
		return nil, err
	}
	diffFiles := map[string]interface{}{}
	for name, wantBytes := range wantFiles {
		gotBytes := gotFiles[name]
		if !bytes.Equal(gotBytes, wantBytes) {
			diffFiles[name] = nil
		}
	}
	for name := range gotFiles {
		if _, ok := wantFiles[name]; !ok {
			diffFiles[name] = nil
		}
	}

	var output []string

	for v := range diffFiles {
		if strings.HasSuffix(v, "_test.go") ||
			strings.HasSuffix(v, "artisanally_handcrafted_code.go") {
			continue
		}
		output = append(output, v)
	}
	return output, nil
}

func Generate(routesPath, outputPath, docsPath string, fs afero.Fs) error {
	err := codegen.Generate(routesPath, outputPath, fs)
	if err != nil {
		return err
	}
	err = docgen.WriteDocs(routesPath, docsPath, fs)
	if err != nil {
		return err
	}

	return nil
}
