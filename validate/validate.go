package validate

import (
	"strings"
)

func Validate(s string) bool {
	if len(s) == 0 {
		return false
	}

	//resp := request.Retrieve(s)
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Couldn't validate the address !")
	//}

	//if strings.Contains(string(body), "href=\"http://") || strings.Contains(string(body), "href=\"https://") {
	//	return true
	//}
	if strings.Contains(s, "href=\"http://") || strings.Contains(string(s), "href=\"https://") {
		return true
	}

	return false
}
