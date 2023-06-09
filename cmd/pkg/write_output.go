package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"pkgsplit/cmd/utils/fileutils"
)

func (p *Pkg) WriteOutput() error {
	outputDir := "output/" + p.Name
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
