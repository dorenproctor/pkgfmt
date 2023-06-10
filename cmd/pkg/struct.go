package pkg

import (
	"go/ast"
	"pkgfmt/cmd/fn"
	"pkgfmt/cmd/impt"
)

type Pkg struct {
	Name     string
	FilePath string
	Body     string
	Fns      []fn.Fn
	Imports  []impt.Impt
	Ast      *ast.File `json:"-"`
}
