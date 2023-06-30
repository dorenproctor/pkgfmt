package cmd

import (
	"fmt"
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
