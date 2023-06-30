package example

import (
	"fmt"
	str "strings"
)

// Does X things
func X(m JustAStringMap) {
	fmt.Println(m, Yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}
