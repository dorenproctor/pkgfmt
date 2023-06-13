package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func (p *Pkg) LoadFromAst() {
	p.VarImports = map[string]impt.Impt{}
	p.StructsIntfs = []PkgPart{}
	var lastIdent *ast.Ident
	ast.Inspect(p.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.Ident:
			lastIdent = n
		case *ast.InterfaceType:
			p.AddStructsIntfs(n, lastIdent)
		case *ast.StructType:
			p.AddStructsIntfs(n, lastIdent)
		case *ast.FuncDecl:
			p.AddFn(n)
		case ast.Decl:
			p.AddVar(n)
		}
		return true
	})
}
