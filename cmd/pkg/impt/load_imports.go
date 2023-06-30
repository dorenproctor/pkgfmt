package impt

import (
	"go/ast"
	"strings"
)

func LoadImports(file *ast.File) []Impt {
	imports := []Impt{}
	for _, importSpec := range file.Imports {
		i := Impt{FullName: importSpec.Path.Value}
		// remove quotes
		i.Name = strings.ReplaceAll(i.FullName, "\"", "")
		// only keep after last /
		i.Name = i.Name[strings.LastIndex(i.Name, "/")+1:]
		if importSpec.Name != nil {
			i.Alias = importSpec.Name.Name
		}
		imports = append(imports, i)
	}
	return imports
}
