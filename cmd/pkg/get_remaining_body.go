package pkg

import (
	"go/parser"
	"go/token"
	"regexp"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func (p *Pkg) GetRemainingBody() (string, error) {
	re := regexp.MustCompile(`(?s)(package\s+\w+\n+)*.*?(import\s+\(?[\s\S]+?\)?)\n\n`)
	body := re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", p.GetBodyWithoutFns(), parser.ParseComments)
	if err != nil {
		return "", err
	}
	imports := impt.GetUsedImports(p.Imports, astFile)
	return fileutils.GetGoFileOutput(p.Name, body, imports), nil
}
