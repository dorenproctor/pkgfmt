package pkg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
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
			p.Fns = append(p.Fns, Fn{
				Name: fn.Name.Name,
				Body: p.Body[fn.Pos()-1 : fn.End()],
				LPos: int(fn.Pos() - 1),
				RPos: int(fn.End()),
			})
		}
	}
}
