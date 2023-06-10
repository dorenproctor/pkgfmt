package fn

import (
	"pkgfmt/cmd/impt"
)

type Fn struct {
	Name        string
	Body        string
	LPos        int
	RPos        int
	Imports     []impt.Impt
	PackageName string
}
