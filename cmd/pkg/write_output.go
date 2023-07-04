package pkg

import (
	"os"

	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

func (p Package) WriteOutput() error {
	outputDir := p.GetOutputDir()
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	if err := pkgpart.WriteFns(outputDir, p.Name, p.Fns); err != nil {
		return err
	}
	if err := pkgpart.WritePkgParts(outputDir+"/types.go", p.Name, p.TypeSpecs); err != nil {
		return err
	}
	if err := pkgpart.WritePkgParts(outputDir+"/vars.go", p.Name, p.Vars); err != nil {
		return err
	}
	// if err := fileutils.OutputGoFile(outputDir+"/comments.go", p.Name, p.GetLooseComments(), nil); err != nil {
	// 	return err
	// }
	if err := p.WriteComments(); err != nil {
		return err
	}
	return nil
}
