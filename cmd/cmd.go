package cmd

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/pkg"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

func Run() {
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgfmt takes one arg, the filepath to a go file or a dir containing go files")
		handleError(e)
	}
	if os.Args[1] == "-v" {
		printVersion()
		return
	}
	p, err := pkg.NewPackage(os.Args[1])
	handleError(err)
	// fmt.Println(p)
	fmt.Println(p.GetAllImports())
	// fmt.Println(fileutils.GetGoFiles(os.Args[1]))
	// p, err := pkg.New(os.Args[1])
	// handleError(err)
	// handleError(p.WriteOutput())
	// fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
