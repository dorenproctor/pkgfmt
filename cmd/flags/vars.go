package flags

var (
	outputDir string
	verbose   bool
	flags     = map[string]any{
		"output-dir": &outputDir,
		"verbose":    &verbose,
	}
)
