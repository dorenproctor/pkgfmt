package pkgpart

import "github.com/dorenproctor/pkgfmt/cmd/pkg/impt"

type PkgPart struct {
	// func/struct/interface/var/const
	Type string
	// name for funcs - blank for others
	Name string
	// the content of this part of the package
	Body string
	// imports used by this package part
	Imports []impt.Impt
}
