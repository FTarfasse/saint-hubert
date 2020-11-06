package validate

import "strings"

func Validate(s string) bool {
	if len(s) == 0 {
		return false
	}
	if strings.Contains(s, "href=\"http://") {
		return true
	}
	if strings.Contains(s, "href=\"https://") {
		return true
	}

	return false
}
