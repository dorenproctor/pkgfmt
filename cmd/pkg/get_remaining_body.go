package pkg

import (
	"pkgsplit/cmd/fn"
	"pkgsplit/cmd/impt"
	"regexp"
)

func (p *Pkg) GetRemainingBody() string {
	// remove package and imports from top of file
	re := regexp.MustCompile(`(?s)(package\s+\w+).*?(import\s+\([\s\S]+?\))\n\n`)
	body := re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	imports := impt.GetUsedImports(p.Imports, body)
	// pretend this is a Fn so we can use its file output
	f := fn.Fn{PackageName: p.Name, Imports: imports, Body: body}
	return f.GetFnFileOutput()
}
