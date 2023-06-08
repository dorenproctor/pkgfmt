package main

import (
	"fmt"
	"pkgsplit/internal/pkg"
)

func main() {
	filePath := "testinput/example/x.go"
	p, err := pkg.New(filePath)
	handleError(err)
	handleError(p.WriteOutput())
	// fmt.Println("updated files in ./output/example/")
	fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
