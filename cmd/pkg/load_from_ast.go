package pkg

import (
	"go/ast"
)

func (p *Pkg) LoadFromAst() {
	p.StructsIntfs = []StructsIntfs{}
	var lastIdent *ast.Ident
	ast.Inspect(p.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.Ident:
			lastIdent = n
		case *ast.InterfaceType:
			p.StructsIntfs = append(p.StructsIntfs, p.NewStructsIntfs(n, lastIdent))
		case *ast.StructType:
			p.StructsIntfs = append(p.StructsIntfs, p.NewStructsIntfs(n, lastIdent))
		case *ast.FuncDecl:
			p.Fns = append(p.Fns, p.NewFn(n))
		}
		return true
	})
}
