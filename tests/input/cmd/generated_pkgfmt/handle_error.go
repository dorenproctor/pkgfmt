package cmd

import (
	"fmt"
	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
