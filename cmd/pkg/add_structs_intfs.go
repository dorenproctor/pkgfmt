package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) AddStructsIntfs(node ast.Node, lastIdent *ast.Ident) {
	body := p.Body[lastIdent.Pos()-6 : node.End()]
	p.StructsIntfs = append(p.StructsIntfs, PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
