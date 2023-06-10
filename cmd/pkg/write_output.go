package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pkgfmt/cmd/utils/fileutils"
	"pkgfmt/cmd/utils/strutils"
)

func (p *Pkg) WriteOutput() error {
	outputDir := fmt.Sprintf("%s/generated_pkgfmt", filepath.Dir(p.FilePath))
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, f := range p.Fns {
		filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(f.Name))
		fileutils.OutputGoFile(filePath, p.Name, f.Body, f.Imports)
		if err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s.go", outputDir, strutils.SnakeCase(p.Name))
	s, err := p.GetRemainingBody()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, []byte(s), 0644)

	// if s != "" {
	// 	return ioutil.WriteFile(filePath, []byte(s), 0644)
	// }
	// return nil
}
