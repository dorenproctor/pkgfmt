package main

import (
	"fmt"
	"os"
	"pkgsplit/cmd/pkg"
)

func main() {
	if len(os.Args) == 1 {
		e := fmt.Errorf("pkgsplit takes one arg, the filepath to a go file")
		handleError(e)
	}
	p, err := pkg.New(os.Args[1])
	handleError(err)
	handleError(p.WriteOutput())

}

func handleError(err error) {
	if err != nil {
		fmt.Println(red(err.Error()))
		os.Exit(1)
	}
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}
