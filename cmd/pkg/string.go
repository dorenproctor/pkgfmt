package pkg

import "github.com/dorenproctor/pkgfmt/cmd/utils/strutils"

func (p Package) String() string {
	return strutils.PrettyJson(p)
}
