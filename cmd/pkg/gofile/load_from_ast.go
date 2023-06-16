package gofile

import (
	"go/ast"
)

func (gf *GoFile) loadFromAst() {
	gf.TypeSpecs = []PkgPart{}
	ast.Inspect(gf.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.TypeSpec:
			gf.addTypeSpec(n)
		case *ast.FuncDecl:
			gf.addFn(n)
		case ast.Decl:
			gf.addVar(n)
		}
		return true
	})
}
