package impt

type Impt struct {
	// Full import path declared in import,
	// eg "go/ast"
	FullName string

	// Name of imported package, after last / of import path,
	// eg "ast" from `import "go/ast"`
	Name string

	// Optional alias variable before the import,
	// eg myalias from `import myalias "go/ast"`
	Alias string
}
