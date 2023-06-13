package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func (p *Pkg) WriteOutput() error {
	outputDir := fmt.Sprintf("%s/generated_pkgfmt", filepath.Dir(p.FilePath))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	if err := p.WriteFns(outputDir); err != nil {
		return err
	}
	if err := p.WriteStructsIntfs(outputDir); err != nil {
		return err
	}
	if err := p.WriteVars(outputDir); err != nil {
		return err
	}
	return nil
	// return p.WriteRemainingBody(outputDir)
}
