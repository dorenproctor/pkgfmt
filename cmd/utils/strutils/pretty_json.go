package strutils

import "encoding/json"

func PrettyJson(a any) string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}
