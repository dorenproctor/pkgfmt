package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

func WritePkgParts(outputFilePath, packageName string, parts []PkgPart) error {
	if len(parts) == 0 {
		return nil
	}
	s := ""
	for _, x := range parts {
		s += x.Body + "\n"
	}
	impts := []impt.Impt{}
	for _, x := range parts {
		impts = append(impts, x.Imports...)
	}
	impts = sliceutils.RemoveDuplicates[impt.Impt](impts)
	err := fileutils.OutputGoFile(outputFilePath, packageName, s, impts)
	if err != nil {
		return err
	}
	return nil
}
