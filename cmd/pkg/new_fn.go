package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) NewFn(node *ast.FuncDecl) PkgPart {
	lpos := int(node.Pos() - 1)
	if node.Doc != nil && len(node.Doc.List) > 0 {
		lpos = int(node.Doc.List[0].Pos() - 1)
	}
	rpos := int(node.End())
	body := p.Body[lpos:rpos]
	return PkgPart{
		Type:    "func",
		Name:    node.Name.Name,
		Body:    body,
		LPos:    lpos,
		RPos:    rpos,
		Imports: impt.GetUsedImports(p.Imports, node),
	}
}
