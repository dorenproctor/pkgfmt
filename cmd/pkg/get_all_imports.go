package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

func (p Package) GetAllImports() []impt.Impt {
	imports := []impt.Impt{}
	for _, f := range p.Files {
		imports = append(imports, f.Imports...)
	}
	return sliceutils.RemoveDuplicates[impt.Impt](imports)
}
