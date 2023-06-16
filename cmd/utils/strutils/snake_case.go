package strutils

import (
	"regexp"
	"strings"
)

// Converts any case into snake_case
func SnakeCase(text string) string {
	snake := regexp.MustCompile("(.)([A-Z][a-z]+)").ReplaceAllString(text, "${1}_${2}")
	snake = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(snake, "${1}_${2}")
	snake = strings.ToLower(snake)
	return snake
}
