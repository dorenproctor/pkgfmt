package fileutils

import (
	"fmt"
	"pkgfmt/cmd/impt"
	"strings"
)

func GetGoFileOutput(packageName, body string, imports []impt.Impt) string {
	body = strings.Trim(body, " \t\n") + "\n"
	return fmt.Sprintf(`package %s
%s
%s`, packageName, impt.GetImportStr(imports), body)
}
