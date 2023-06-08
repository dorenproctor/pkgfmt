package pkg

import (
	"go/ast"
	"pkgsplit/internal/fn"
	"pkgsplit/internal/impt"
)

type Pkg struct {
	Name    string      `json:"name"`
	Fns     []fn.Fn     `json:"fns"`
	Body    string      `json:"body"`
	Ast     *ast.File   `json:"-"`
	Imports []impt.Impt `json:"imports"`
}
