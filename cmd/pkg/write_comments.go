package pkg

import (
	"strings"

	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func (p Package) WriteComments() error {
	comments := []string{}
	for _, gf := range p.Files {
		comments = append(comments, gf.LooseComments...)
	}
	if len(comments) == 0 {
		return nil
	}
	s := strings.Join(comments, "\n")
	return fileutils.OutputGoFile(
		p.GetOutputDir()+"/comments.go", p.Name, s, nil)
}
