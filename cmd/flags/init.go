package flags

import "flag"

// runs once when flags package is first imported
func init() {
	flag.StringVar(&OutputDir, "output-dir", "", "")
	flag.StringVar(&OutputDir, "o", "", "")
	flag.BoolVar(&Verbose, "verbose", false, "")
	flag.BoolVar(&Verbose, "v", false, "")
}
