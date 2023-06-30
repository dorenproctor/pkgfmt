package cmd

import (
	"fmt"
	"runtime"
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
