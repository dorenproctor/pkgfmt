package main

import (
	"fmt"
	"pkgsplit/cmd/pkg"
)

func main() {
	filePath := "testinput/example/x.go"
	p, err := pkg.New(filePath)
	handleError(err)
	handleError(p.WriteOutput())
	// p.ExtractImports()
	// fmt.Println("updated files in ./output/example/")
	fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
