package cmd

import (
	"os"
	"fmt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg"
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
	handleError(p.WriteOutput())
}
