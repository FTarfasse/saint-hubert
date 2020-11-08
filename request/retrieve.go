package request

import (
	"net/http"
)

//func Retrieve(url string) *http.Response {
func Retrieve(url string) *http.Response {

	resp, err := http.Get(url)
	if err != nil {
		panic("Couldn't retrieve the address")
	}
	//b, _ := ioutil.ReadAll(resp.Body)
	// body := resp.Body
	// formatted, _ := ioutil.ReadAll(body)
	// string := string(formatted)
	// fmt.Printf("Resp body is %s\n", string)
	// defer resp.Body.Close()

	return resp
}
