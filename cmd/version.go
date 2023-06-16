package cmd

import "fmt"

var Version string

func printVersion() {
	if Version == "" {
		Version = "unpublished version"
	}
	fmt.Println("pkgfmt", Version)
}
