package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) NewVar(node ast.Node) {
	lpos := int(node.Pos() - 1)
	rpos := int(node.End())
	body := p.Body[lpos:rpos]
	varType := ""
	if strings.HasPrefix(body, "var") {
		varType = "var"
	} else if strings.HasPrefix(body, "const") {
		varType = "const"
	}
	if varType == "" {
		return
	}
	p.Vars = append(p.Vars, PkgPart{
		Type:    varType,
		LPos:    lpos,
		RPos:    rpos,
		Body:    body,
		Imports: impt.GetUsedImports(p.Imports, node),
	})
}
