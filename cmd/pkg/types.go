package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

// package
type Pkg struct {
	// name of package
	Name string
	// path to file that was read
	FilePath string
	// full package contents
	Body string
	// list of the body of all var/const declarations
	Vars []PkgPart
	// imports used in var declarations
	VarImports map[string]impt.Impt
	// each func will get its own file
	Fns []PkgPart
	// all imports used by this package
	Imports []impt.Impt
	// things declared with "type"
	TypeSpecs []PkgPart
	// abstract syntax tree for the package
	Ast *ast.File `json:"-"`
}

type PkgPart struct {
	// func/struct/interface/var/const
	Type string
	// Name for funcs - blank for others
	Name string
	// the content of this part of the package
	Body string
	// imports used by this package part
	Imports []impt.Impt
}
