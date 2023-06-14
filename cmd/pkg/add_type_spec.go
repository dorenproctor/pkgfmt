package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) AddTypeSpec(node ast.Node) {
	lpos := strings.LastIndex(p.Body[0:node.Pos()], "type ")
	if lpos < 0 {
		return
	}
	body := p.Body[lpos:node.End()]
	p.TypeSpecs = append(p.TypeSpecs, PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
