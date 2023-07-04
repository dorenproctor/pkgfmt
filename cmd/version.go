package cmd

import (
	"fmt"
	"runtime"

	"github.com/dorenproctor/pkgfmt/cmd/flags"
)

const (
	toolName    = "pkgfmt"
	notProvided = "[not provided]"
)

var (
	version   = notProvided
	buildDate = notProvided
	gitCommit = notProvided
	platform  = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

func printVersion() {
	if flags.Verbose {
		fmt.Printf("Version:        %s\n", version)
		fmt.Printf("BuildDate:      %s\n", buildDate)
		fmt.Printf("GitCommit:      %s\n", gitCommit)
		fmt.Printf("Platform:       %s\n", platform)
		fmt.Printf("GoVersion:      %s\n", runtime.Version())
		fmt.Printf("Compiler:       %s\n", runtime.Compiler)
	} else {
		fmt.Printf("%s %s (%s)\n", toolName, version, buildDate)
	}
}
