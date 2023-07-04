package flags

var (
	// default is <arg>/generated_pkgfmt
	OutputDir string
	// only used for version, print more info, default false
	Verbose bool
	// prevents settings flags multiple times which causes error
	alreadySetFlags bool
)
