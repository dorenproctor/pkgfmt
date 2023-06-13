package pkg

import (
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func (p *Pkg) WriteVars(outputDir string) error {
	s := strings.Join(p.Vars, "\n")
	filePath := outputDir + "/vars.go"
	err := fileutils.OutputGoFile(filePath, p.Name, s, []impt.Impt{})
	if err != nil {
		return err
	}
	return nil
}
