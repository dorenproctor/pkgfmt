package pkg

import (
	"go/parser"
	"go/token"
	"pkgsplit/cmd/fn"
	"pkgsplit/cmd/impt"
	"regexp"
)

func (p *Pkg) GetRemainingBody() (string, error) {
	// remove package and imports from top of file
	re := regexp.MustCompile(`(?s)(package\s+\w+).*?(import\s+\([\s\S]+?\))\n\n`)
	body := re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", p.GetBodyWithoutFns(), parser.ParseComments)
	if err != nil {
		return "", err
	}
	imports := impt.GetUsedImports(p.Imports, astFile)
	// imports := impt.GetUsedImportsStr(p.Imports, body)
	// pretend this is a Fn so we can use its file output
	f := fn.Fn{PackageName: p.Name, Imports: imports, Body: body}
	return f.GetFnFileOutput(), nil
}
