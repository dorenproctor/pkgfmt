package pkgpart

import (
	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

// takes multiple Fns and returns a list of all Impts used without duplicates
func GetImportSetUnion(fns []PkgPart) []impt.Impt {
	impts := []impt.Impt{}
	for _, fn := range fns {
		impts = append(impts, fn.Imports...)
	}
	return sliceutils.RemoveDuplicates[impt.Impt](impts)
}
