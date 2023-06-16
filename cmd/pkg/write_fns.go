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
