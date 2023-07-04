package flags

import (
	"flag"
	"os"
)

// instead of flag.Parse, use the first arg as either filepath or 'version'
// and read flags from 2nd arg onward
func ParseFlags() {
	if len(os.Args) > 2 {
		flag.CommandLine.Parse(os.Args[2:])
	}
}
