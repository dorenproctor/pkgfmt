package example

import (
	"fmt"
)

var (
	s string
	t = "ttt"
)

const foo = "bar"

// Does X things
func x() {
	fmt.Println(yzzy("x!"))
}

// yzzy man
// what more is there to say?
func yzzy(s string) string {
	return s
}
