package types

import (
	"go/ast"
	"github.com/dorenproctor/github.com/dorenproctor/pkgfmt/types/cmd/impt"
)

// Package
type types.Pkg struct {
	Name         string
	FilePath     string
	Body         string
	Fns          []Fn
	Imports      []impt.Impt
	StructOrIntf []StructOrIntf
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
type StructOrIntf struct {
	Name string
	Body string
	LPos int
	RPos int
}
