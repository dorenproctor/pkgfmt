package pkg

import "go/ast"

func (p *Pkg) NewStructOrIntf(n ast.Node, lastIdent *ast.Ident) StructOrIntf {
	lpos := int(lastIdent.Pos() - 6)
	rpos := int(n.End())
	return StructOrIntf{
		LPos: lpos,
		RPos: rpos,
		Body: p.Body[lpos:rpos],
	}
}
