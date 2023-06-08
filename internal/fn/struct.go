package fn

import (
	"pkgsplit/internal/impt"
)

type Fn struct {
	Name        string      `json:"name"`
	Body        string      `json:"body"`
	LPos        int         `json:"lpos"`
	RPos        int         `json:"rpos"`
	Imports     []impt.Impt `json:"imports"`
	PackageName string      `json:"packagename"`
}
