package pkg

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

func GetPackage(filePath string) (Package, error) {
	p := Package{}
	if err := p.LoadBody(filePath); err != nil {
		return p, err
	}
	if err := p.LoadAst(filePath); err != nil {
		return p, err
	}
	p.LoadFns()
	return p, nil
}

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
	p.ast = file
	if p.ast.Name != nil {
		p.Name = p.ast.Name.Name
	}
	return nil
}

func (p *Package) LoadFns() {
	for _, decl := range p.ast.Decls {
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

func (p Package) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "failed to parse Package struct"
	}
	return string(jsonData)
}

type Package struct {
	Name string `json:"name"`
	Fns  []Fn   `json:"fns"`
	Body string `json:"body"`
	ast  *ast.File
}

type Fn struct {
	Name string `json:"name"`
	Body string `json:"body"`
	LPos int    `json:"lpos"`
	RPos int    `json:"rpos"`
}
