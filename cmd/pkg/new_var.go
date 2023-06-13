package pkg

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) NewVar(node ast.Node) {
	s := p.Body[int(node.Pos()-1):int(node.End())]
	if strings.HasPrefix(s, "var") || strings.HasPrefix(s, "const") {
		p.Vars = append(p.Vars, s)
		for _, i := range impt.GetUsedImports(p.Imports, node) {
			name := i.Alias
			if name == "" {
				name = i.Name
			}
			p.VarImports[name] = i
		}
	}
}
