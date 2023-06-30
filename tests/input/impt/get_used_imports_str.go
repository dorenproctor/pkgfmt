package impt

import "strings"

// Import
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

// Get imports referenced in a string - not reliable
func GetUsedImportsStr(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if importUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}

func importUsed(i Impt, body string) bool {
	if i.Alias != "" {
		if strings.Contains(body, i.Alias+".") {
			return true
		}
	} else {
		if strings.Contains(body, i.Name+".") {
			return true
		}
	}
	return false
}
