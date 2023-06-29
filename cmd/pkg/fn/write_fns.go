package fn

import (
	"fmt"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

func WriteFns(outputDir, packageName string, fns []gofile.PkgPart) error {
	nameToFns := GroupDuplicatesFns(fns)
	for name, fns := range nameToFns {
		body := ""
		for _, f := range fns {
			body += fmt.Sprintf("\n%s\n", f.Body)
		}
		filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(name))
		err := fileutils.OutputGoFile(filePath, packageName, body, GetImportSetUnion(fns))
		if err != nil {
			return err
		}
	}
	return nil
}
