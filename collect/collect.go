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

// func removeDuplicate(array []Result, i int) {
//
// 	// Remove the element at i i from a.
// 	copy(array[i:], array[i+1:]) // Shift a[i+1:] left one i.
// 	array[len(array)-1] = ""     // Erase last element (write zero value).
// 	array = array[:len(array)-1] // Truncate slice.
//
// 	//fmt.Println(a) // [A B D E]
// }
//
// func findDuplicate(res []Result, i int) bool {
// 	for k := 0; r := range res; k++ {
// 		if res[r].Address == val {
// 			return true
// 		}
// 	}
// 	return false
// }

func Find(res []Result, search string) bool {
	r := len(res)
	count := 0
	for i := 0; i < r; i++ {
		if res[i].Address == search {
			count++
		}
	}

	return count != 1
}