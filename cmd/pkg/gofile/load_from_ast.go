package gofile

import (
	"go/ast"
)

func (gf *GoFile) loadFromAst() {
	// go/ast does not store CommentGroups on TypeSpecs for some reason
	var lastCommentGroup *ast.CommentGroup
	// track which comments are found inside Inspect to identify loose comments
	isDocstring := map[*ast.CommentGroup]bool{}
	ast.Inspect(gf.Ast, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.CommentGroup:
			isDocstring[n] = true
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
	for _, comment := range gf.Ast.Comments {
		if !isDocstring[comment] {
			s := gf.Body[comment.Pos()-1 : comment.End()]
			gf.LooseComments = append(gf.LooseComments, s)
			// gf.LooseComments = append(gf.LooseComments, "// "+comment.Text())
		}
	}
}
