package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/sliceutils"
)

func (p *Pkg) WriteStructsIntfs(outputDir string) error {
	if len(p.StructsIntfs) == 0 {
		return nil
	}
	s := ""
	for _, x := range p.StructsIntfs {
		s += x.Body + "\n"
	}
	filePath := outputDir + "/types.go"
	impts := []impt.Impt{}
	for _, x := range p.StructsIntfs {
		impts = append(impts, x.Imports...)
	}
	impts = sliceutils.RemoveDuplicates[impt.Impt](impts)
	err := fileutils.OutputGoFile(filePath, p.Name, s, impts)
	if err != nil {
		return err
	}
	return nil
}
