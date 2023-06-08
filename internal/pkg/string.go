package pkg

import "encoding/json"

func (p Package) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "failed to parse Package struct"
	}
	return string(jsonData)
}
