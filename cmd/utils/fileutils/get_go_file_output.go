package fileutils

import (
	"fmt"
	"pkgsplit/cmd/impt"
)

func GetGoFileOutput(packageName, body string, imports []impt.Impt) string {
	return fmt.Sprintf(`package %s
%s
%s`, packageName, impt.GetImportStr(imports), body)
}
