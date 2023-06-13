package pkg

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"regexp"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

func (p *Pkg) WriteRemainingBody(outputDir string) error {
	filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(p.Name))
	re := regexp.MustCompile(`(?s)(package\s+\w+\n+)*.*?(import\s+\(?[\s\S]+?\)?)\n\n`)
	body := re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", p.GetBodyWithoutFns(), parser.ParseComments)
	if err != nil {
		return err
	}
	imports := impt.GetUsedImports(p.Imports, astFile)
	s, err := fileutils.GetGoFileOutput(p.Name, body, imports), nil
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, []byte(s), 0644)
}
