package pkg

import (
	"go/ast"
)

func (p *Pkg) LoadFromAst() {
	p.StructOrIntf = []StructOrIntf{}
	var lastIdent *ast.Ident
	ast.Inspect(p.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.Ident:
			lastIdent = n
		case *ast.InterfaceType:
			p.StructOrIntf = append(p.StructOrIntf, p.NewStructOrIntf(n, lastIdent))
		case *ast.StructType:
			p.StructOrIntf = append(p.StructOrIntf, p.NewStructOrIntf(n, lastIdent))
		case *ast.FuncDecl:
			p.Fns = append(p.Fns, p.NewFn(n))
		}
		return true
	})
}
