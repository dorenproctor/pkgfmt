package fn

import (
	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

// takes multiple Fns and returns a list of all Impts used without duplicates
func GetImportSetUnion(fns []gofile.PkgPart) []impt.Impt {
	impts := []impt.Impt{}
	for _, fn := range fns {
		impts = append(impts, fn.Imports...)
	}
	return sliceutils.RemoveDuplicates[impt.Impt](impts)
}
