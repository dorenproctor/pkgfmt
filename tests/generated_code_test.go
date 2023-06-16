package tests

import (
	"os"
	"path"
	"testing"

	"github.com/dorenproctor/pkgfmt/cmd/pkg"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedCode(t *testing.T) {
	t.Parallel()
	inputFiles := []string{
		"input/example/example.go",
		"input/impt/get_used_imports_str.go",
	}
	for _, inputFile := range inputFiles {
		// declaring out here fixes closure loop issue with t.Parallel()
		fileName := inputFile
		t.Run(fileName, func(t *testing.T) {
			t.Parallel()
			inputDir := path.Dir(fileName)
			p, err := pkg.NewPackage(fileName)
			assert.NoError(t, err)
			assert.NoError(t, os.RemoveAll(inputDir+"/generated_pkgfmt"))
			assert.NoError(t, p.WriteOutput())
			if os.Getenv("OVERWRITE_TEST_EXPECTED_OUTPUT") == "true" {
				assert.NoError(t, fileutils.CopyFilesInDir(inputDir+"/generated_pkgfmt", inputDir+"/expected_output"))
			}
			assert.NoError(t, fileutils.Diff(inputDir+"/generated_pkgfmt", inputDir+"/expected_output"))
		})
	}
}
