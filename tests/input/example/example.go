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

type Person struct {
	Name   string
	secret string
	fmt.Formatter
	fmt.Scanner
}

type Asdf interface{}

// Does X things
func x() {
	fmt.Println(yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}

// yzzy man
// what more is there to say?
func yzzy(s string) string {
	return str.Replace(s, "yzzy", "", -1)
}
