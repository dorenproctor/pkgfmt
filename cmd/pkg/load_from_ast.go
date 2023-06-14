package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) LoadFromAst() {
	p.VarImports = map[string]impt.Impt{}
	p.TypeSpecs = []PkgPart{}
	ast.Inspect(p.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.TypeSpec:
			p.AddTypeSpec(n)
		case *ast.FuncDecl:
			p.AddFn(n)
		case ast.Decl:
			p.AddVar(n)
		}
		return true
	})
}
