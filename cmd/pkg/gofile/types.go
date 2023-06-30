package gofile

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

type GoFile struct {
	// path to file that was read
	FilePath string
	// name of package
	PackageName string
	// full file contents
	Body string
	// all funcs in file
	Fns []pkgpart.PkgPart
	// var and const declarations
	Vars []pkgpart.PkgPart
	// all imports used in file
	Imports []impt.Impt
	// all things declared with "type"
	TypeSpecs []pkgpart.PkgPart
	// abstract syntax tree for the package
	Ast *ast.File `json:"-"`
}
