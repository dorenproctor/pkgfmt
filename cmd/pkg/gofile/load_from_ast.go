package gofile

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

func (gf *GoFile) loadFromAst() {
	gf.TypeSpecs = []pkgpart.PkgPart{}
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
