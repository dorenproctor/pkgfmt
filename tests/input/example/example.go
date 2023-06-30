package example

import (
	"fmt"
	astalias "go/ast"
	str "strings"
)

var (
	// this is not lost
	S string
	T = fmt.Sprintf("%s %s", "", "")
)

const Foo = "bar"

// this is currently lost....
type Fooer func(string) string
type JustAStringMap map[string]string

// this is currently lost....
type Person struct {
	// this is not lost
	Name   string
	secret string
	fmt.Formatter
	fmt.Scanner
	Fooer
}

// shh
func (p Person) TellSecret() string {
	return p.secret
}

type Asdf interface{}

// Does X things
func X(m JustAStringMap) {
	fmt.Println(astalias.ArrayType{}, m, Yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}

// yzzy man
// what more is there to say?
func Yzzy(s string) string {
	return str.Replace(s, "yzzy", "", -1)
}

// loose comments
// should loose comments go somewhere?
