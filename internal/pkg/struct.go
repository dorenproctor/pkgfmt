package pkg

import (
	"go/ast"
	"pkgsplit/internal/fn"
	"pkgsplit/internal/impt"
)

type Pkg struct {
	Name    string
	Fns     []fn.Fn
	Body    string
	Imports []impt.Impt
	Ast     *ast.File `json:"-"`
}
