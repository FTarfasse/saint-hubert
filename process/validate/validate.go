package validate

import (
	"strings"
)

// Validate function verifies if the current string contains valid http links
func Validate(s string) bool {
	if len(s) == 0 {
		return false
	}

	if strings.Contains(s, "href=\"http://") || strings.Contains(string(s), "href=\"https://") {
		return true
	}

	return false
}
