package cmd

import (
	"fmt"
	"os"
	"runtime"
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
	if verboseVersion() {
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

func verboseVersion() bool {
	return len(os.Args) > 2 && (os.Args[2] == "-v" || os.Args[2] == "--verbose")
}
