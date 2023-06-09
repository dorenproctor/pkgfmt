package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"pkgsplit/cmd/utils/fileutils"
)

func (p *Pkg) WriteOutput() error {
	// outputDir := "output/" + p.Name
	// outputDir := filepath.Dir(p.FilePath)
	outputDir := fmt.Sprintf("%s/generated_pkgsplit", filepath.Dir(p.FilePath))
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, f := range p.Fns {
		filePath := fmt.Sprintf("%s/%s.go", outputDir, f.Name)
		fileutils.OutputGoFile(filePath, f.PackageName, f.Body, f.Imports)
		if err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s.go", outputDir, p.Name)
	s, err := p.GetRemainingBody()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, []byte(s), 0644)
}
