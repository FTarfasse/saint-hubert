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
	Code    int
	Source  string
}

//const ErrNoLinksFound = errors.New("No links at this address !")

//var datas []Result

func Collect(url string) (datas []Result, err error) {

	r := request.Retrieve(url)
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	b := string(rb)
	links, err := slice.SliceLinks(b)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		resp := request.Retrieve(link)
		status := resp.Status
		code := resp.StatusCode
		result := Result{Address: link, Status: status, Code: code, Source: url}
		datas = append(datas, result)
	}

	return datas, err
}

func Squeeze(array []Result) []Result {
	var squeezed []Result
	length := len(array)

	if CheckDoubloons(array) == true {
		for i := 0; i < length; i++ {
			_, times := Count(array[i:], array[i].Address)
			if times == 1 {
				squeezed = append(squeezed, array[i])
			}
		}
	}

	return squeezed
}

func CheckDoubloons(arr []Result) bool {
	for _, a := range arr {
		_, count := Count(arr, a.Address)
		if count != 1 {
			return true
		}
	}

	return false
}

func Count(res []Result, search string) (bool, int) {
	r := len(res)
	count := 0
	for i := 0; i < r; i++ {
		if res[i].Address == search {
			count++
		}
	}

	return count == 1, count
}
