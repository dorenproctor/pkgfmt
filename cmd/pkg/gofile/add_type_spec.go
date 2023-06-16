package gofile

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (gf *GoFile) AddTypeSpec(node ast.Node) {
	lpos := strings.LastIndex(gf.Body[0:node.Pos()], "type ")
	if lpos < 0 {
		return
	}
	body := gf.Body[lpos:node.End()]
	gf.TypeSpecs = append(gf.TypeSpecs, PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(gf.Imports, node),
	})
}
