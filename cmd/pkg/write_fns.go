package pkg

import (
	"fmt"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

func WriteFns(outputDir, packageName string, fns []gofile.PkgPart) error {
	for _, f := range fns {
		filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(f.Name))
		err := fileutils.OutputGoFile(filePath, packageName, f.Body, f.Imports)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Pkg) WriteFns(outputDir string) error {
	for _, f := range p.Fns {
		filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(f.Name))
		err := fileutils.OutputGoFile(filePath, p.Name, f.Body, f.Imports)
		if err != nil {
			return err
		}
	}
	return nil
}
