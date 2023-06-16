package strutils

import "encoding/json"

// Convert to JSON string with formatting and indentation
//
// If JSON conversion fails, returns "{}"
func PrettyJson(a any) string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}
