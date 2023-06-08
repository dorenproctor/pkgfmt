package main

import (
	"fmt"
	"pkgsplit/internal/pkg"
)

func main() {
	// Specify the file path of the Go file to parse
	filePath := "testinput/example/x.go"
	p, err := pkg.New(filePath)
	handleError(err)
	fmt.Println(p)
	handleError(p.WriteOutput())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
