package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/pkgpart"
)

type Package struct {
	// Package name taken from first file
	Name string
	// path that was passed as argument
	FilePath string
	// either all go files in dir passed or single length if file passed
	Files []gofile.GoFile
	// all imports used in package
	Imports []impt.Impt
	// all funcs in package
	Fns []pkgpart.PkgPart
	// all var and const declarations in package
	Vars []pkgpart.PkgPart
	// all things declared with "type" in package
	TypeSpecs []pkgpart.PkgPart
}
