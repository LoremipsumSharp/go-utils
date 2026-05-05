package string

import "strings"

func isNilOrEmpty(text *string) bool {
	return len(*text) == 0
}

func IsNilOrWhitespace(s *string) bool {
	return isNilOrEmpty(s) || len(strings.TrimSpace(*s)) == 0
}

func IsEmptyOrWhitespace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
