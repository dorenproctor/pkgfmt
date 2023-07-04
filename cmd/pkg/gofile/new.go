package gofile

import (
	"go/parser"
	"go/token"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

func New(filePath string) (GoFile, error) {
	gf := GoFile{
		FilePath: filePath,
		Fns:      []pkgpart.PkgPart{},
		Vars:     []pkgpart.PkgPart{},
		Imports:  []impt.Impt{},
	}
	// get body
	fileBytes, err := os.ReadFile(filePath)
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
	// load all imports used in file
	gf.Imports = impt.LoadImports(gf.Ast)
	// load TypeSpecs, Fns, and Vars
	gf.loadFromAst()
	return gf, nil
}
