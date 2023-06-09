package pkg

import (
	"go/ast"
	"pkgsplit/cmd/fn"
	"pkgsplit/cmd/impt"
)

type Pkg struct {
	Name     string
	FilePath string
	Body     string
	Fns      []fn.Fn
	Imports  []impt.Impt
	Ast      *ast.File `json:"-"`
}
