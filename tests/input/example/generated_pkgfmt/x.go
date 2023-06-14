package example

import (
	"fmt"
	str "strings"
)

// Does X things
func x(m JustAStringMap) {
	fmt.Println(m, yzzy(str.Trim("x!", str.ToLower("ASDF"))))
}
