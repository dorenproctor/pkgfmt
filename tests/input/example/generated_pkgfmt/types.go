package example

import (
	"fmt"
)

type Fooer func(string) string

type JustAStringMap map[string]string

type Person struct {
	Name   string
	secret string
	fmt.Formatter
	fmt.Scanner
	Fooer
}

type Asdf interface{}
