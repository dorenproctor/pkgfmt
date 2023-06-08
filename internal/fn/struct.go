package fn

import (
	"pkgsplit/internal/impt"
)

type Fn struct {
	Name        string
	Body        string
	LPos        int
	RPos        int
	Imports     []impt.Impt
	PackageName string
}
