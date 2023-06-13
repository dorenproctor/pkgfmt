package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) NewStructsIntfs(node ast.Node, lastIdent *ast.Ident) PkgPart {
	lpos := int(lastIdent.Pos() - 6)
	rpos := int(node.End())
	body := p.Body[lpos:rpos]
	return PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	}
}
