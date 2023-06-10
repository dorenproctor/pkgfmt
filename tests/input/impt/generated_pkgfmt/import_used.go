package impt

import (
	"strings"
)

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
