package pkg

import (
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func New(filePath string) (*Pkg, error) {
	p := Pkg{FilePath: filePath}
	// get body
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &p, err
	}
	p.Body = string(fileBytes)
	// get ast
	fset := token.NewFileSet()
	p.Ast, err = parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return &p, err
	}
	// get package name
	if p.Ast.Name != nil {
		p.Name = p.Ast.Name.Name
	}
	// load all imports used in file
	p.Imports = impt.LoadImports(p.Ast)
	// load StructsIntfs, Fns, and Vars
	p.LoadFromAst()
	return &p, nil
}
