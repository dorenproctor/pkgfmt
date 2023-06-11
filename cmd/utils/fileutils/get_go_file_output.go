package fileutils

import (
	"fmt"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func GetGoFileOutput(packageName, body string, imports []impt.Impt) string {
	body = strings.Trim(body, " \t\n") + "\n"
	return fmt.Sprintf(`package %s
%s
%s`, packageName, impt.GetImportStr(imports), body)
}
