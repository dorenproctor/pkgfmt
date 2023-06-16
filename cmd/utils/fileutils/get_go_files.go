package fileutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetGoFiles(filePath string) ([]string, error) {
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
			goFiles = append(goFiles, s)
		}
	}
	if info.IsDir() {
		files, err := ioutil.ReadDir(filePath)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			addIfGoFile(filePath + "/" + f.Name())
		}
	} else {
		addIfGoFile(filePath)
	}
	return goFiles, nil

}
