package gofile

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
)

func (gf *GoFile) addVar(node ast.Node) {
	body := gf.Body[node.Pos()-1 : node.End()]
	varType := ""
	if strings.HasPrefix(body, "var") {
		varType = "var"
	} else if strings.HasPrefix(body, "const") {
		varType = "const"
	}
	if varType == "" {
		return
	}
	gf.Vars = append(gf.Vars, PkgPart{
		Type:    varType,
		Body:    body,
		Imports: impt.GetUsedImports(gf.Imports, node),
	})
}
