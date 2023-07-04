package gofile

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

func (gf *GoFile) loadFromAst() {
	gf.TypeSpecs = []pkgpart.PkgPart{}
	// go/ast does not store CommentGroups on TypeSpecs for some reason
	var lastCommentGroup *ast.CommentGroup
	ast.Inspect(gf.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.CommentGroup:
			lastCommentGroup = n
		case *ast.TypeSpec:
			gf.addTypeSpec(n, lastCommentGroup)
			lastCommentGroup = nil
		case *ast.FuncDecl:
			gf.addFn(n)
			lastCommentGroup = nil
		case ast.Decl:
			gf.addVar(n)
			lastCommentGroup = nil
		}
		return true
	})
}
