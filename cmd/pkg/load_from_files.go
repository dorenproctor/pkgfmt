package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

func (p *Package) loadFromFiles() {
	// load imports
	imports := []impt.Impt{}
	for _, f := range p.Files {
		imports = append(imports, f.Imports...)
	}
	p.Imports = sliceutils.RemoveDuplicates[impt.Impt](imports)
	// load vars
	vars := []gofile.PkgPart{}
	for _, f := range p.Files {
		vars = append(vars, f.Vars...)
	}
	p.Vars = vars
	// load fns
	fns := []gofile.PkgPart{}
	for _, f := range p.Files {
		fns = append(fns, f.Fns...)
	}
	p.Fns = fns
}
