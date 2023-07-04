package gofile

import (
	"go/ast"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

func (gf *GoFile) addTypeSpec(node ast.Node, lastCommentGroup *ast.CommentGroup) {
	lpos := strings.LastIndex(gf.Body[0:node.Pos()], "type ")
	if lpos < 0 {
		return
	}
	if lastCommentGroup != nil && len(lastCommentGroup.List) > 0 {
		lpos = int(lastCommentGroup.List[0].Pos()) - 1
	}
	body := gf.Body[lpos:node.End()]
	gf.TypeSpecs = append(gf.TypeSpecs, pkgpart.PkgPart{
		Type:    strings.Split(body, " ")[2],
		Body:    body,
		Imports: impt.GetUsedImports(gf.Imports, node),
	})
}
