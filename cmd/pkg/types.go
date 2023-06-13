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
	Vars []string
	// imports used in var declarations
	VarImports map[string]impt.Impt
	// each func will get its own file
	Fns []Fn
	// all imports used by this package
	Imports []impt.Impt
	// structs or interfaces
	StructsIntfs []StructsIntfs
	// abstract syntax tree for the package
	Ast *ast.File `json:"-"`
}

// func
type Fn struct {
	Name    string
	Body    string
	LPos    int
	RPos    int
	Imports []impt.Impt
}

// interface
type StructsIntfs struct {
	Name string
	Body string
	LPos int
	RPos int
}
