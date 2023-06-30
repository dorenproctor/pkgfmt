package cmd

import (
	"os"
)

func verboseVersion() bool {
	return len(os.Args) > 2 && (os.Args[2] == "-v" || os.Args[2] == "--verbose")
}
