package pkg

import (
	"os"
)

func (p Package) WriteOutput() error {
	outputDir := p.GetOutputDir()
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
