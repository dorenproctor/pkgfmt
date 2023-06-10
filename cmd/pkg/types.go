package pkg

import (
	"go/ast"
	"pkgfmt/cmd/impt"
)

// Package
type Pkg struct {
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
