package fileutils

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/impt"
)

func OutputGoFile(filePath, packageName, body string, imports []impt.Impt) error {
	s := goFileOutput(packageName, body, imports)
	return ioutil.WriteFile(filePath, []byte(s), 0644)
}

func goFileOutput(packageName, body string, imports []impt.Impt) string {
	body = strings.Trim(body, " \t\n") + "\n"
	return fmt.Sprintf(`package %s
%s
%s`, packageName, impt.GetImportStr(imports), body)
}
