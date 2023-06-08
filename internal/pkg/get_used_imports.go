package pkg

import (
	"strings"
)

func GetUsedImports(imports []Import, body string) []Import {
	used := []Import{}
	for _, i := range imports {
		if strings.Contains(body, i.Name+".") {
			used = append(used, i)
		}
	}
	return used
}
