package pkg

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/flags"
)

func (p Package) GetOutputDir() string {
	outputDirFlag := flags.GetFlag[string]("output-dir")
	if *outputDirFlag != "" {
		return *outputDirFlag
	}
	filePath := p.FilePath
	if strings.HasSuffix(filePath, ".go") {
		filePath = filepath.Dir(p.FilePath)
	}
	return fmt.Sprintf("%s/generated_pkgfmt", filePath)
}
