package main

import (
	"fmt"
	"os"
	"pkgfmt/cmd/pkg"
	"pkgfmt/cmd/utils/strutils"
)

func main() {
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgfmt takes one arg, the filepath to a go file")
		handleError(e)
	}
	p, err := pkg.New(os.Args[1])
	handleError(err)
	handleError(p.WriteOutput())
	// fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
