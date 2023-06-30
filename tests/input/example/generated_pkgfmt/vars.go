package example

import (
	"fmt"
)

var (
	// this is not lost
	S string
	T = fmt.Sprintf("%s %s", "", "")
)

const Foo = "bar"
