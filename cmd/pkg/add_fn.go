package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) AddFn(node *ast.FuncDecl) {
	lpos := node.Pos() - 1
	// include docs if they're there
	if node.Doc != nil && len(node.Doc.List) > 0 {
		lpos = node.Doc.List[0].Pos() - 1
	}
	body := p.Body[lpos:node.End()]
	p.Fns = append(p.Fns, PkgPart{
		Type:    "func",
		Name:    node.Name.Name,
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
