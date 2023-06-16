package gofile

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
)

type GoFile struct {
	// path to file that was read
	FilePath string
	// name of package
	PackageName string
	// full file contents
	Body string
	// each func will get its own file
	Fns []PkgPart
	// list of the body of all var/const declarations
	Vars []PkgPart
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
	// name for funcs - blank for others
	Name string
	// the content of this part of the package
	Body string
	// imports used by this package part
	Imports []impt.Impt
}
