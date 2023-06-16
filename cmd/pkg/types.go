package pkg

import (
	"github.com/dorenproctor/pkgfmt/cmd/pkg/gofile"
	"github.com/dorenproctor/pkgfmt/cmd/pkg/impt"
)

type Package struct {
	// Package name taken from first file
	Name string
	// path to file that was read
	FilePath string
	// data about either the file passed or all go files in dir passed
	Files []gofile.GoFile
	// all imports used by this package
	Imports []impt.Impt
	// each func will get its own file
	Fns []gofile.PkgPart
	// list of the body of all var/const declarations
	Vars []gofile.PkgPart
	// things declared with "type"
	TypeSpecs []gofile.PkgPart
}
