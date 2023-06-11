package load

import (
	"go/parser"
	"go/token"
)

func (p *types.Pkg) LoadAst(filePath string) error {
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
