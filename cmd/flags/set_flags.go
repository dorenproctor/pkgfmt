package flags

import (
	"flag"
	"os"
)

func SetFlags() {
	if alreadySetFlags {
		return
	}
	alreadySetFlags = true
	flag.StringVar(&OutputDir, "output-dir", "", "")
	flag.StringVar(&OutputDir, "o", "", "")
	flag.BoolVar(&Verbose, "verbose", false, "")
	flag.BoolVar(&Verbose, "v", false, "")
	// instead of flag.Parse, use the first arg as either filepath or 'version'
	// and read flags from 2nd arg onward
	if len(os.Args) > 2 {
		flag.CommandLine.Parse(os.Args[2:])
	}
}
