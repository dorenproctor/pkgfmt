package example

import (
	"fmt"
)

type Person struct {
	Name   string
	secret string
	fmt.Formatter
	fmt.Scanner
}

type Asdf interface{}
