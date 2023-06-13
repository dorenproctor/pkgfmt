package pkg

import "go/ast"

func (p *Pkg) NewStructsIntfs(node ast.Node, lastIdent *ast.Ident) StructsIntfs {
	lpos := int(lastIdent.Pos() - 6)
	rpos := int(node.End())
	return StructsIntfs{
		LPos: lpos,
		RPos: rpos,
		Body: p.Body[lpos:rpos],
	}
}
