package example

import (
	"fmt"
	str "strings"
)

var (
	s string
	t = fmt.Sprintf("%s %s", "", "")
)

const foo = "bar"

type Fooer func(string) string
type JustAStringMap map[string]string

type Person struct {
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
func x(m JustAStringMap) {
	fmt.Println(m, yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}

// yzzy man
// what more is there to say?
func yzzy(s string) string {
	return str.Replace(s, "yzzy", "", -1)
}

// loose comments
// should loose comments go somewhere?
