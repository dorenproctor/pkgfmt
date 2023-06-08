package main

import (
	"fmt"
	"pkgsplit/internal/pkg"
)

func main() {
	// Specify the file path of the Go file to parse
	filePath := "testinput/example/x.go"
	p, err := pkg.GetPackage(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
