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
	if os.Args[1] == "version" {
		printVersion()
		return
	}
	p, err := pkg.NewPackage(os.Args[1])
	handleError(err)
	fmt.Println(p)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
