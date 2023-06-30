package example

import (
	"fmt"
	astalias "go/ast"
	str "strings"
)

// Does X things
func X(m JustAStringMap) {
	fmt.Println(astalias.ArrayType{}, m, Yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}
