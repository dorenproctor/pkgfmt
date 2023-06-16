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
		"input/errors",
	}
	for _, inputFile := range inputFiles {
		// declaring out here fixes closure loop issue with t.Parallel()
		fileName := inputFile
		t.Run(fileName, func(t *testing.T) {
			t.Parallel()
			// inputDir := path.Dir(fileName)
			p, err := pkg.NewPackage(fileName)
			outputDir := p.GetOutputDir()
			expectedDir := path.Dir(outputDir) + "/expected_output"
			assert.NoError(t, err)
			assert.NoError(t, os.RemoveAll(outputDir))
			assert.NoError(t, p.WriteOutput())
			if os.Getenv("OVERWRITE_TEST_EXPECTED_OUTPUT") == "true" {
				assert.NoError(t, fileutils.CopyFilesInDir(outputDir, expectedDir))
			}
			assert.NoError(t, fileutils.Diff(outputDir, expectedDir))
		})
	}
}
