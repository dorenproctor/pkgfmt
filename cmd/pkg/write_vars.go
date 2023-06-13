package pkg

import (
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func (p *Pkg) WriteVars(outputDir string) error {
	if len(p.Vars) == 0 {
		return nil
	}
	filePath := outputDir + "/vars.go"
	varDeclarations := strings.Join(p.Vars, "\n")
	impts := []impt.Impt{}
	for _, i := range p.VarImports {
		impts = append(impts, i)
	}
	err := fileutils.OutputGoFile(filePath, p.Name, varDeclarations, impts)
	if err != nil {
		return err
	}
	return nil
}
