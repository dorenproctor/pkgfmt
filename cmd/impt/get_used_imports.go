package impt

import (
	"fmt"
	"go/ast"
)

func GetUsedImportsFn(imports []Impt, f *ast.FuncDecl) []Impt {
	alreadyUsed := map[string]bool{}
	used := []Impt{}
	ast.Inspect(f, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SelectorExpr:
			for _, i := range imports {
				name := fmt.Sprintf("%s", n.X)
				if name == i.Alias || name == i.Name {
					if !alreadyUsed[name] {
						used = append(used, i)
						alreadyUsed[name] = true
					}
				}
			}
		}
		return true
	})
	return used
}

func GetUsedImportsStr(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if ImportUsed(i, body) {
			used = append(used, i)
		}
	}
	return used
}
