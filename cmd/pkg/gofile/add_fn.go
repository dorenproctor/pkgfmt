package gofile

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (gf *GoFile) addFn(node *ast.FuncDecl) {
	lpos := node.Pos() - 1
	// include docs if they're there
	if node.Doc != nil && len(node.Doc.List) > 0 {
		lpos = node.Doc.List[0].Pos() - 1
	}
	body := gf.Body[lpos:node.End()]
	gf.Fns = append(gf.Fns, PkgPart{
		Type:    "func",
		Name:    node.Name.Name,
		Body:    body,
		Imports: impt.GetUsedImports(gf.Imports, node),
	})
}
