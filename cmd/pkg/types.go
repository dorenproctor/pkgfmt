package pkg

import (
	"go/ast"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

// Package
type Pkg struct {
	Name         string
	FilePath     string
	Body         string
	Vars         []string
	Fns          []Fn
	Imports      []impt.Impt
	StructsIntfs []StructsIntfs
	Ast          *ast.File `json:"-"`
}

// Function
type Fn struct {
	Name    string
	Body    string
	LPos    int
	RPos    int
	Imports []impt.Impt
}

// Interface
type StructsIntfs struct {
	Name string
	Body string
	LPos int
	RPos int
}
