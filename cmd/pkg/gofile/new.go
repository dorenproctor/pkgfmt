package gofile

import (
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func New(filePath string) (GoFile, error) {
	gf := GoFile{
		FilePath: filePath,
		Fns:      []PkgPart{},
		Vars:     []PkgPart{},
		Imports:  []impt.Impt{},
	}
	// get body
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return gf, err
	}
	gf.Body = string(fileBytes)
	// get ast
	fset := token.NewFileSet()
	gf.Ast, err = parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return gf, err
	}
	// get package name
	if gf.Ast.Name != nil {
		gf.PackageName = gf.Ast.Name.Name
	}
	gf.loadFromAst()
	// load all imports used in file
	gf.Imports = impt.LoadImports(gf.Ast)
	return gf, nil
}
