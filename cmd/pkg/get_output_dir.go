package pkg

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (p Package) GetOutputDir() string {
	filePath := p.FilePath
	if strings.HasSuffix(filePath, ".go") {
		filePath = filepath.Dir(p.FilePath)
	}
	return fmt.Sprintf("%s/generated_pkgfmt", filePath)
}
