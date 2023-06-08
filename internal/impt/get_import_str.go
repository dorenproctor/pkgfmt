package impt

import (
	"fmt"
	"strings"
)

func GetImportStr(imports []Impt) string {
	output := ""
	if len(imports) > 0 {
		importStrs := []string{}
		for _, i := range imports {
			importStrs = append(importStrs, i.String())
		}

		output = fmt.Sprintf(
			"\nimport (\n\t%s\n)\n",
			strings.Join(importStrs, "\n\t"),
		)
	}
	return output
}
