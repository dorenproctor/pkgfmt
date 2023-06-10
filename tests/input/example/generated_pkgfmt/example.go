package example

import (
	"fmt"
)

var (
	s string
	t = fmt.Sprintf("%s %s", "", "")
)

const foo = "bar"

type Person struct {
	Name   string
	secret string
}

type Asdf interface{}
