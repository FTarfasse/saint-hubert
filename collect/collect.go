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

func Squeeze(array []Result) []Result {
	var newArr []Result

	if CheckDoubloons(array) == true {

		for i := 0; i < len(array); i++ {
			_, times := AddressPresentOnce(array, array[i].Address)
			if times > 1 {
				array = append(array[:i], array[i+1:]...) // remove from array
			}
			if times == 1 {
				newArr = append(newArr, array[i]) // add to squeezed array
			}
		}

	}

	return array
}

func CheckDoubloons(arr []Result) bool {
	for _, a := range arr {
		_, count := AddressPresentOnce(arr, a.Address)
		for count != 1 {
			return true
		}
	}
	return false
}

func AddressPresentOnce(res []Result, search string) (bool, int) {
	r := len(res)
	count := 0
	for i := 0; i < r; i++ {
		if res[i].Address == search {
			count++
		}
	}
	return count == 1, count
}
