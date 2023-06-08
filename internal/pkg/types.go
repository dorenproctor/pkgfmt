package pkg

import (
	"go/ast"
)

type Package struct {
	Name    string    `json:"name"`
	Fns     []Fn      `json:"fns"`
	Body    string    `json:"body"`
	Ast     *ast.File `json:"-"`
	Imports []Import  `json:"imports"`
}

type Fn struct {
	Name    string   `json:"name"`
	Body    string   `json:"body"`
	LPos    int      `json:"lpos"`
	RPos    int      `json:"rpos"`
	Imports []Import `json:"imports"`
	pkg     Package
}

type Import struct {
	Name           string `json:"name"`
	NameWithQuotes string `json:"qname"`
	Alias          string `json:"alias"`
}
