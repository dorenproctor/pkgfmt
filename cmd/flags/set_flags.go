package flags

import (
	"flag"
	"os"
)

func SetFlags() {
	flag.StringVar(&outputDir, "output-dir", "", "")
	flag.StringVar(&outputDir, "o", "", "")
	flag.BoolVar(&verbose, "verbose", false, "")
	flag.BoolVar(&verbose, "v", false, "")
	// instead of flag.Parse, use the first arg as either filepath or 'version'
	// and read flags from 2nd arg onward
	if len(os.Args) > 2 {
		flag.CommandLine.Parse(os.Args[2:])
	}
}
