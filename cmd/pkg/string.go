package pkg

import "github.com/dorenproctor/pkgfmt/cmd/utils/strutils"

func (p Pkg) String() string {
	return strutils.PrettyJson(p)
}
