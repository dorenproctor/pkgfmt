package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func (p Package) WriteOutput() error {
	outputDir := fmt.Sprintf("%s/generated_pkgfmt", filepath.Dir(p.FilePath))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	if err := WriteFns(outputDir, p.Name, p.Fns); err != nil {
		return err
	}
	if err := WritePkgParts(outputDir+"/types.go", p.Name, p.TypeSpecs); err != nil {
		return err
	}
	if err := WritePkgParts(outputDir+"/vars.go", p.Name, p.Vars); err != nil {
		return err
	}
	return nil
}

func (p *Pkg) WriteOutput() error {
	outputDir := fmt.Sprintf("%s/generated_pkgfmt", filepath.Dir(p.FilePath))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	if err := p.WriteFns(outputDir); err != nil {
		return err
	}
	if err := p.WritePkgParts(outputDir+"/types.go", p.TypeSpecs); err != nil {
		return err
	}
	if err := p.WritePkgParts(outputDir+"/vars.go", p.Vars); err != nil {
		return err
	}
	return nil
}
