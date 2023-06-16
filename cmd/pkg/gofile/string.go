package gofile

import "encoding/json"

func (gf GoFile) String() string {
	jsonData, err := json.MarshalIndent(gf, "", "  ")
	if err != nil {
		return "failed to parse Package struct"
	}
	return string(jsonData)
}
