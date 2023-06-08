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

// Does X things
func x() {
	fmt.Println(yzzy(str.Trim("x!", "")))
}

// yzzy man
// what more is there to say?
func yzzy(s string) string {
	return str.Replace(s, "yzzy", "", -1)
}
