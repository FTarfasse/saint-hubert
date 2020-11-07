package request

import (
	"io/ioutil"
	"net/http"
)

type Result struct {
	Address string
	Status  string
	Body    string
}

func Retrieve(url string) Result {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	status := resp.Status

	result := Result{Address: url, Status: status, Body: string(body)}

	return result
}
