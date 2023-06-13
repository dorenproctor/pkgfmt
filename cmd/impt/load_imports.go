package impt

import (
	"go/ast"
	"strings"
)

func LoadImports(file *ast.File) []Impt {
	imports := []Impt{}
	for _, importSpec := range file.Imports {
		i := Impt{Name: importSpec.Path.Value}
		i.NameWithQuotes = importSpec.Path.Value
		i.Name = strings.Replace(importSpec.Path.Value, "\"", "", 2)
		if importSpec.Name != nil {
			i.Alias = importSpec.Name.Name
		}
		imports = append(imports, i)
	}
	return imports
}
