package pkg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

func (p *Package) LoadBody(filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	p.Body = string(fileBytes)
	return nil
}

func (p *Package) LoadAst(filePath string) error {
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

func (p *Package) LoadFns() {
	for _, decl := range p.Ast.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			// for _, x := range fn.Doc.List {
			// 	// fmt.Println(x.)
			// }
			lpos := int(fn.Pos() - 1)
			if len(fn.Doc.List) > 0 {
				lpos = int(fn.Doc.List[0].Pos() - 1)
			}
			rpos := int(fn.End())
			p.Fns = append(p.Fns, Fn{
				Name: fn.Name.Name,
				Body: p.Body[lpos:rpos],
				LPos: lpos,
				RPos: rpos,
			})
		}
	}
}
func (p *Package) LoadImports() {
	p.Imports = []string{}
	for _, importSpec := range p.Ast.Imports {
		s := strings.Trim(importSpec.Path.Value, "\"")
		p.Imports = append(p.Imports, s)
	}
}
