package cmd

import (
	"fmt"
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/utils/strutils"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(strutils.Red(err.Error()))
		os.Exit(1)
	}
}
