package pkg

import (
	"go/parser"
	"go/token"
	"pkgfmt/cmd/impt"
	"pkgfmt/cmd/utils/fileutils"
	"regexp"
)

func (p *Pkg) GetRemainingBody() (string, error) {
	// ... does this work?
	// remove package and imports from top of file
	re := regexp.MustCompile(`(?s)(package\s+\w+\n+)*.*?(import\s+\([\s\S]+?\))\n\n`)
	body := re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	// re = regexp.MustCompile(fmt.Sprintf(`(\b%s\b|\s)`, "package"))
	// body = re.ReplaceAllString(p.GetBodyWithoutFns(), "")
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "", p.GetBodyWithoutFns(), parser.ParseComments)
	if err != nil {
		return "", err
	}
	imports := impt.GetUsedImports(p.Imports, astFile)
	return fileutils.GetGoFileOutput(p.Name, body, imports), nil

	// s := fileutirls.GetGoFileOutput(p.Name, body, imports)
	// s2 := strings.Trim(strings.Replace(s, "package "+p.Name, "", -1), " \t\n")
	// fmt.Println(s)
	// fmt.Println("===============")
	// fmt.Println(s2)
	// fmt.Println("===============")
	// fmt.Println(p.GetBodyWithoutFns())
	// if s2 == "" {
	// 	return "", nil
	// }
	// return s, nil
}
