package cmd

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

// if err not nil print it in red and exit 1
func HandleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
