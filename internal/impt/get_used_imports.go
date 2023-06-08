package impt

import (
	"strings"
)

func GetUsedImports(imports []Impt, body string) []Impt {
	used := []Impt{}
	for _, i := range imports {
		if strings.Contains(body, i.Name+".") {
			used = append(used, i)
		}
	}
	return used
}
