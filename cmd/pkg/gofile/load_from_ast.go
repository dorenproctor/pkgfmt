package gofile

import (
	"go/ast"
)

func (gf *GoFile) LoadFromAst() {
	gf.TypeSpecs = []PkgPart{}
	ast.Inspect(gf.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.TypeSpec:
			gf.AddTypeSpec(n)
		case *ast.FuncDecl:
			gf.AddFn(n)
		case ast.Decl:
			gf.AddVar(n)
		}
		return true
	})
}
