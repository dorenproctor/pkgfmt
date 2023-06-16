package main

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

var Version string

func main() {
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgfmt takes one arg, the filepath to a go file or a dir containing go files")
		handleError(e)
	}
	if os.Args[1] == "-v" {
		printVersion()
		return
	}
	fmt.Println(fileutils.GetGoFiles(os.Args[1]))
	// p, err := pkg.New(os.Args[1])
	// handleError(err)
	// handleError(p.WriteOutput())
	// fmt.Println(p)
}

func printVersion() {
	if Version == "" {
		Version = "unpublished version"
	}
	fmt.Println("pkfmt", Version)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
