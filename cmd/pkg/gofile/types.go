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
	// each func will get its own file
	Fns []pkgpart.PkgPart
	// list of the body of all var/const declarations
	Vars []pkgpart.PkgPart
	// all imports used by this package
	Imports []impt.Impt
	// things declared with "type"
	TypeSpecs []pkgpart.PkgPart
	// abstract syntax tree for the package
	Ast *ast.File `json:"-"`
}
