package collect

import (
	request "../request"
	slice "../slicelinks"
	"io/ioutil"
	"log"
	"errors"
)

type Result struct {
	Address string
	Status  string
	Code    int
	Source  string
}

var ErrNoLinksFound = errors.New("No links at this address !")

func Collect(url string) (datas []Result, err error) {

	r := request.Retrieve(url)
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//b := string(rb)
	links, err := slice.SliceLinks(string(rb))
	if err != nil {
		return datas, ErrNoLinksFound
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

func Squeeze(array []Result, issuesOnly bool) (datas []Result) {
	var squeezed []Result
	doubles := CheckDoubloons(array)


	if doubles == false && issuesOnly == false {
		return array
	}

	if doubles == true && issuesOnly == false {
		for i := 0; i < len(array); i++ {
			_, times := Count(array[i:], array[i].Address)
			if times == 1 {
				squeezed = append(squeezed, array[i])
			}
		}
		return squeezed
	}

	if doubles == false && issuesOnly == true {
		onlyIssues := make([]Result,0)
		var x int
		for j := 0; j < len(array); j++ {
			//fmt.Printf("\n \033[0;31m%v\033[0m\n", array[x])
			if array[j].Code > 199 && array[j].Code < 300 {
				// NOTHING
			} else {
				onlyIssues = append(onlyIssues, array[x])
				x++
			}
		}
		return onlyIssues
	}

	if doubles == true && issuesOnly == true {
		var fullFiltering []Result
		for k := 0; k < len(squeezed); k++ {
			if squeezed[k].Code > 199 && squeezed[k].Code < 300 {
				// NOTHING
			} else {
				fullFiltering = append(fullFiltering, squeezed[k])
			}
		}
		return fullFiltering
	}

	return
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
