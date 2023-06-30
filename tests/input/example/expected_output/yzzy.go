package example

import (
	str "strings"
)

// yzzy man
// what more is there to say?
func Yzzy(s string) string {
	return str.Replace(s, "yzzy", "", -1)
}
