package pkg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"pkgsplit/internal/fn"
	"pkgsplit/internal/impt"
	"strings"
)

func (p *Pkg) LoadBody(filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	p.Body = string(fileBytes)
	return nil
}

func (p *Pkg) LoadAst(filePath string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	p.Ast = file
	if p.Ast.Name != nil {
		p.Name = p.Ast.Name.Name
	}
	return nil
}

func (p *Pkg) LoadFns() {
	for _, decl := range p.Ast.Decls {
		if f, ok := decl.(*ast.FuncDecl); ok {
			lpos := int(f.Pos() - 1)
			if len(f.Doc.List) > 0 {
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
				Imports:     impt.GetUsedImports(p.Imports, body),
			})
		}
	}
}

func (p *Pkg) LoadImports() {
	p.Imports = []impt.Impt{}
	for _, importSpec := range p.Ast.Imports {
		i := impt.Impt{Name: importSpec.Path.Value}
		i.NameWithQuotes = importSpec.Path.Value
		i.Name = strings.Replace(importSpec.Path.Value, "\"", "", 2)
		if importSpec.Name != nil {
			i.Alias = importSpec.Name.Name
		}
		p.Imports = append(p.Imports, i)
	}
}
