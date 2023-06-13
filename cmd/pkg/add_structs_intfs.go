package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) AddStructsIntfs(node ast.Node, lastIdent *ast.Ident) {
	lpos := int(lastIdent.Pos() - 6)
	rpos := int(node.End())
	body := p.Body[lpos:rpos]
	p.StructsIntfs = append(p.StructsIntfs, PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
