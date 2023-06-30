package example

import (
	"fmt"
)

var (
	// this is not lost
	s string
	t = fmt.Sprintf("%s %s", "", "")
)

const foo = "bar"
