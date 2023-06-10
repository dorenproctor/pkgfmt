package tests

import (
	"os"
	"path"
	"pkgfmt/cmd/pkg"
	"pkgfmt/cmd/utils/fileutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedCode(t *testing.T) {
	t.Parallel()
	inputFiles := []string{
		"input/example/x.go",
		"input/impt/get_used_imports_str.go",
	}
	for _, fileName := range inputFiles {
		t.Run(fileName, func(t *testing.T) {
			inputDir := path.Dir(fileName)
			p, err := pkg.New(fileName)
			assert.NoError(t, err)
			assert.NoError(t, os.RemoveAll(inputDir+"/generated_pkgfmt"))
			assert.NoError(t, p.WriteOutput())
			// generate expected_output by uncommenting this line
			// assert.NoError(t, fileutils.CopyFilesInDir(inputDir+"/generated_pkgfmt", inputDir+"/expected_output"))
			assert.NoError(t, fileutils.Diff(inputDir+"/generated_pkgfmt", inputDir+"/expected_output"))
		})
	}
}
