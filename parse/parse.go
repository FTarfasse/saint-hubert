package parse

func Parse(s string) string {
	if Validate(s) == false {
		return "No links at this address !"
	}
	return ""
}