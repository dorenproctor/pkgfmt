package example

import (
	"fmt"
)

// does foo things!
type Fooer func(string) string

type JustAStringMap map[string]string

type Person struct {
	// this is not lost
	Name   string
	secret string
	fmt.Formatter
	fmt.Scanner
	Fooer
}

type Asdf interface{}
