package parse

import validate "../validate"

func Parse(s string) string {
	if validate.Validate(s) == false {
		return "No links at this address !"
	}

	// create struct

	return ""
}