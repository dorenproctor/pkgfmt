package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
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
	err := fileutils.OutputGoFile(filePath, p.Name, s, []impt.Impt{})
	if err != nil {
		return err
	}
	return nil
}
