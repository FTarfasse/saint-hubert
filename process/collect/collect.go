// collect package includes the functions to process the HTML string into an exploitable input
package collect

import (
	//"../request"
	slice "../slicelinks"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	t "../types"
)

var ErrNoLinksFound = errors.New("No links at this address !")

// Collect executes the initial HTTP call to retrieve the full HTML string, extract the links and executes the calls to
// each one
func Collect(url string) (datas []t.Result, err error) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	links, err := slice.SliceLinks(string(rb))
	if err != nil {
		return datas, ErrNoLinksFound
	}

	for _, link := range links {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatalln(err)
		}
		status := resp.Status
		code := resp.StatusCode
		result := t.Result{Address: link, Status: status, Code: code, Source: url}
		datas = append(datas, result)
	}

	return datas, err
}

// this function removes duplicate links from the array of Result struct
func Squeeze(array []t.Result, issuesOnly bool) (datas []t.Result) {
	var squeezed []t.Result
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
		onlyIssues := make([]t.Result, 0)
		var x int
		for j := 0; j < len(array); j++ {
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
		var fullFiltering []t.Result
		var fullTreat []t.Result
		// squeeze first
		for b := 0; b < len(array); b++ {
			_, times := Count(array[b:], array[b].Address)
			if times == 1 {
				fullFiltering = append(fullFiltering, array[b])
			}
		}

		for k := 0; k < len(fullFiltering); k++ {
			if fullFiltering[k].Code > 199 && fullFiltering[k].Code < 300 {
				// NOTHING
			} else {
				fullTreat = append(fullTreat, fullFiltering[k])
			}
		}
		return fullTreat
	}

	return
}

// this function checks if the current array has any duplicate link (Result.Address value)
func CheckDoubloons(arr []t.Result) bool {
	for _, a := range arr {
		_, count := Count(arr, a.Address)
		if count != 1 {
			return true
		}
	}

	return false
}

// checks the number of occurrences of the link (Result.Address value)
func Count(res []t.Result, search string) (bool, int) {
	r := len(res)
	count := 0
	for i := 0; i < r; i++ {
		if res[i].Address == search {
			count++
		}
	}

	return count == 1, count
}