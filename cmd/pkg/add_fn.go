package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) AddFn(node *ast.FuncDecl) {
	lpos := int(node.Pos() - 1)
	if node.Doc != nil && len(node.Doc.List) > 0 {
		lpos = int(node.Doc.List[0].Pos() - 1)
	}
	rpos := int(node.End())
	body := p.Body[lpos:rpos]
	p.Fns = append(p.Fns, PkgPart{
		Type:    "func",
		Name:    node.Name.Name,
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
