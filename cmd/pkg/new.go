package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/utils/fileutils"
)

func NewPackage(filePath string) (Package, error) {
	p := Package{FilePath: filePath, Files: []gofile.GoFile{}}
	fileNames, err := fileutils.GetGoFiles(filePath)
	if err != nil {
		return p, err
	}
	for _, f := range fileNames {
		gf, err := gofile.New(f)
		if err != nil {
			return p, err
		}
		p.Files = append(p.Files, gf)
	}
	p.loadFromFiles()
	p.Name = p.Files[0].PackageName
	return p, nil
}
