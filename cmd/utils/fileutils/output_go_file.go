package fileutils

import (
	"io/ioutil"
	"pkgsplit/cmd/impt"
)

func OutputGoFile(filePath, packageName, body string, imports []impt.Impt) error {
	s := GetGoFileOutput(packageName, body, imports)
	return ioutil.WriteFile(filePath, []byte(s), 0644)
}