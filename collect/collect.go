package collect

import (
	slice "../slice"
	validate "../validate"
)

func Collect(s string) string {
	if validate.Validate(s) == false {
		return "No links at this address !"
	}

	links := slice.Slice(s)

	return ""
}
