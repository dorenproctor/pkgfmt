package pkg

import (
	"go/ast"
	"pkgfmt/cmd/impt"
)

type Pkg struct {
	Name     string
	FilePath string
	Body     string
	Fns      []Fn
	Imports  []impt.Impt
	Ast      *ast.File `json:"-"`
}

type Fn struct {
	Name        string
	Body        string
	LPos        int
	RPos        int
	Imports     []impt.Impt
	PackageName string
}
