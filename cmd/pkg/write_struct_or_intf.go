package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func (p *Pkg) WriteStructOrIntf(outputDir string) error {
	s := ""
	for _, x := range p.StructOrIntf {
		s += x.Body + "\n"
	}
	filePath := outputDir + "/types.go"
	err := fileutils.OutputGoFile(filePath, p.Name, s, []impt.Impt{})
	if err != nil {
		return err
	}
	return nil
}
