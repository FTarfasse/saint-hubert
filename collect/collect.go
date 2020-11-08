package collect

import (
	request "../request"
	slice "../slicelinks"
	"io/ioutil"
	"log"
)

type Result struct {
	Address string
	Status  string
	Code int
	Source string
}

var datas []Result

func Collect(url string) []Result {

	r := request.Retrieve(url)
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	b := string(rb)
	links := slice.SliceLinks(b)

	for _, link := range links {
		resp := request.Retrieve(link)
		status := resp.Status
		code := resp.StatusCode
		result := Result{Address: link, Status: status, Code: code, Source: url}
		datas = append(datas, result)
	}

	return datas
}