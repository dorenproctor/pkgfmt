package cmd

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/flags"
	"github.com/dorenproctor/pkgfmt/cmd/pkg"
)

// run pkgfmt
func Run() {
	flags.ParseFlags()
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgfmt takes one arg, the filepath to a go file or a dir containing go files")
		HandleError(e)
	}
	if os.Args[1] == "version" {
		printVersion()
		return
	}
	p, err := pkg.NewPackage(os.Args[1])
	HandleError(err)
	HandleError(p.WriteOutput())
}
