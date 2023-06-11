package load

import (
	"go/ast"
)

func (p *types.Pkg) LoadStructOrIntfs() {
	p.StructOrIntf = []StructOrIntf{}
	var lastIdent *ast.Ident
	ast.Inspect(p.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.Ident:
			lastIdent = n
		case *ast.InterfaceType:
			lpos := int(lastIdent.Pos() - 6)
			rpos := int(n.End())
			p.StructOrIntf = append(p.StructOrIntf, StructOrIntf{
				LPos: lpos,
				RPos: rpos,
				Body: p.Body[lpos:rpos],
			})
		case *ast.StructType:
			lpos := int(lastIdent.Pos() - 6)
			rpos := int(n.End())
			p.StructOrIntf = append(p.StructOrIntf, StructOrIntf{
				LPos: lpos,
				RPos: rpos,
				Body: p.Body[lpos:rpos],
			})
		}
		return true
	})
}
