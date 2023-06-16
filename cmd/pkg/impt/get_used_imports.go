package impt

import (
	"fmt"
	"go/ast"
)

func GetUsedImports(imports []Impt, n ast.Node) []Impt {
	alreadyUsed := map[string]bool{}
	used := []Impt{}
	ast.Inspect(n, func(node ast.Node) bool {
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
