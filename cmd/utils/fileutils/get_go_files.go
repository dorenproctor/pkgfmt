package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// If filePath is a directory, get all files inside it ending in "go".
//
// If filePath is a valid file ending in ".go", return single length list
//
// If filePath is a valid file *not* ending in ".go", return zero length list
//
// If includeTestFiles is false, exclude files ending in "_test.go"
func GetGoFiles(filePath string, includeTestFiles bool) ([]string, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s does not exist", filePath)
		}
		return nil, err
	}
	goFiles := []string{}
	addIfGoFile := func(s string) {
		if strings.HasSuffix(s, ".go") {
			if includeTestFiles || !strings.HasSuffix(s, "_test.go") {
				goFiles = append(goFiles, s)
			}
		}
	}
	if info.IsDir() {
		files, err := ioutil.ReadDir(filePath)
		if err != nil {
			return nil, err
		}
		filePath = strings.TrimSuffix(filePath, "/")
		for _, f := range files {
			addIfGoFile(filePath + "/" + f.Name())
		}
	} else {
		addIfGoFile(filePath)
	}
	return goFiles, nil
}
