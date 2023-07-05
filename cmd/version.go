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
	buildVersion = notProvided
	buildDate    = notProvided
	gitCommit    = notProvided
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

func SetVersion(version, commit, date string) {
	buildVersion = version
	gitCommit = commit
	buildDate = date
}

func printVersion() {
	if verboseVersion() {
		fmt.Printf("Version:        %s\n", buildVersion)
		fmt.Printf("BuildDate:      %s\n", buildDate)
		fmt.Printf("GitCommit:      %s\n", gitCommit)
		fmt.Printf("Platform:       %s\n", platform)
		fmt.Printf("GoVersion:      %s\n", runtime.Version())
		fmt.Printf("Compiler:       %s\n", runtime.Compiler)
	} else {
		fmt.Printf("%s %s (%s)\n", toolName, buildVersion, buildDate)
	}
}

func verboseVersion() bool {
	return len(os.Args) > 2 && (os.Args[2] == "-v" || os.Args[2] == "--verbose")
}
