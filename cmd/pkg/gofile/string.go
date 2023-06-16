package gofile

import "github.com/dorenproctor/pkgfmt/cmd/utils/strutils"

func (gf GoFile) String() string {
	return strutils.PrettyJson(gf)
}
