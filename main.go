package main

import (
	"fmt"
	"pkgsplit/cmd/pkg"
)

func main() {
	filePath := "tests/input/example/x.go"
	p, err := pkg.New(filePath)
	handleError(err)
	// handleError(p.WriteOutput())
	fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
