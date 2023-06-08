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
		if f, ok := decl.(*ast.FuncDecl); ok {
			lpos := int(f.Pos() - 1)
			if len(f.Doc.List) > 0 {
				lpos = int(f.Doc.List[0].Pos() - 1)
			}
			rpos := int(f.End())
			body := p.Body[lpos:rpos]
			p.Fns = append(p.Fns, Fn{
				Name:    f.Name.Name,
				Body:    body,
				LPos:    lpos,
				RPos:    rpos,
				pkg:     *p,
				Imports: GetUsedImports(p.Imports, body),
			})
		}
	}
}

func (p *Package) LoadImports() {
	p.Imports = []Import{}
	for _, importSpec := range p.Ast.Imports {
		i := Import{Name: importSpec.Path.Value}
		i.NameWithQuotes = importSpec.Path.Value
		i.Name = strings.Replace(importSpec.Path.Value, "\"", "", 2)
		if importSpec.Name != nil {
			i.Alias = importSpec.Name.Name
		}
		p.Imports = append(p.Imports, i)
	}
}

// func (p *Package) LoadImports() {
// 	p.Imports = []string{}
// 	for _, importSpec := range p.Ast.Imports {
// 		s := addQuotes(strings.Trim(importSpec.Path.Value, "\""))
// 		if importSpec.Name != nil {
// 			s = fmt.Sprintf("%s %s", importSpec.Name.Name, s)
// 		}
// 		p.Imports = append(p.Imports, s)
// 	}
// }

// func addQuotes(s string) string {
// 	return fmt.Sprintf("\"%s\"", s)
// }
