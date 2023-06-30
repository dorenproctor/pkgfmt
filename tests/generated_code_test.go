package tests

import (
	"os"
	"path"
	"testing"

	"github.com/dorenproctor/pkgfmt/cmd"
	"github.com/dorenproctor/pkgfmt/cmd/pkg"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
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
			p, err := pkg.NewPackage(fileName)
			outputDir := p.GetOutputDir()
			expectedDir := path.Dir(outputDir) + "/expected_output"
			assertNoError(t, err)
			assertNoError(t, os.RemoveAll(outputDir))
			os.Args = []string{"", fileName}
			cmd.Run()
			if os.Getenv("OVERWRITE_TEST_EXPECTED_OUTPUT") == "true" {
				assertNoError(t, fileutils.CopyFilesInDir(outputDir, expectedDir))
			}
			assertNoError(t, fileutils.Diff(outputDir, expectedDir))
		})
	}
}

func assertNoError(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}
