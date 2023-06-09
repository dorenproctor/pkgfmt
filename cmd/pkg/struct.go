package pkg

import (
	"go/ast"
	"pkgsplit/cmd/fn"
	"pkgsplit/cmd/impt"
)

type Pkg struct {
	Name    string
	Fns     []fn.Fn
	Body    string
	Imports []impt.Impt
	Ast     *ast.File `json:"-"`
}
