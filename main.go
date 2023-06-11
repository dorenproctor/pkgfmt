package main

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/pkg"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

var Version = "unofficial version"

func main() {
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgfmt takes one arg, the filepath to a go file")
		handleError(e)
	}
	if os.Args[1] == "-v" {
		fmt.Println("pkfmt", Version)
		return
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
