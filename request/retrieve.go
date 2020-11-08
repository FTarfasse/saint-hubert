package request

import (
	"net/http"
)

func Retrieve(url string) *http.Response {

	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't retrieve the address")
	}

	return resp
}