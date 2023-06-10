package tests

import (
	"os"
	"path"
	"pkgfmt/cmd/pkg"
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
			// assert.NoError(t, fileutils.CopyDir(inputDir+"/generated_pkgfmt", inputDir+"/expected_output"))
			// diff, err := fileutils.Diff(inputDir+"/generated_pkgfmt", inputDir+"/expected_output")
			// assert.NoError(t, err)
			// fmt.Println(diff)
		})
	}
}

// for _, tc := range GeneratedCodeTestCases {
// 	t.Run(tc.FileName, func(t *testing.T) {
// 		// t.Parallel()

// 		inputFile := fmt.Sprintf("%s/%s", tc.InputDir, tc.FileName)
// 		p, err := pkg.New(inputFile)
// 		assert.NoError(t, err)

// 		err = os.RemoveAll(fmt.Sprintf("%s/generated_pkgfmt", tc.InputDir))
// 		assert.NoError(t, err)

// 		assert.NoError(t, p.WriteOutput())
// 	})
// }
// }

// type GeneratedCodeTestCase struct {
// 	InputDir string
// 	FileName string
// }

// var GeneratedCodeTestCases = []GeneratedCodeTestCase{
// 	{
// 		InputDir: "input/example",
// 		FileName: "x.go",
// 	},
// }
