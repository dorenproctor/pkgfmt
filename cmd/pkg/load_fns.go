package pkg

import (
	"go/ast"
	"pkgfmt/cmd/fn"
	"pkgfmt/cmd/impt"
)

func (p *Pkg) LoadFns() {
	for _, decl := range p.Ast.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			lpos := int(f.Pos() - 1)
			if f.Doc != nil && len(f.Doc.List) > 0 {
				lpos = int(f.Doc.List[0].Pos() - 1)
			}
			rpos := int(f.End())
			body := p.Body[lpos:rpos]
			p.Fns = append(p.Fns, fn.Fn{
				Name:        f.Name.Name,
				Body:        body,
				LPos:        lpos,
				RPos:        rpos,
				PackageName: p.Name,
				Imports:     impt.GetUsedImports(p.Imports, f),
			})
		}
	}
}
