// collect package includes the functions to process the HTML string into an exploitable input
package collect

import (
	//"../request"
	slice "../slicelinks"
	types "../types"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var ErrNoLinksFound = errors.New("No links at this address !")

// TODO: refactor into several functions
// Collect executes the initial HTTP call to retrieve the full HTML string, extract the links and executes the calls to
// each one
func Collect(url string) (datas []types.Result, err error) {
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

	var wg sync.WaitGroup
	resultChannel := make(chan types.Result)

	for _, link := range links {
		wg.Add(1)
		go func(l string, channel chan types.Result, wg *sync.WaitGroup) {
			// TODO: find better way to handle unreachable host
			resp, err := http.Get(l)
			if err != nil {
				result := types.Result{Address: l, Status: "Unreachable host", Code: 0, Source: url}
				channel <- result
			}
			if err == nil {
				result := types.Result{Address: l, Status: resp.Status, Code: resp.StatusCode, Source: url}
				channel <- result
			}
			wg.Done()
		}(link, resultChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for res := range resultChannel {
		datas = append(datas, res)
	}

	return datas, err
}

// this function removes duplicate links from the array of Result struct
func Squeeze(array []types.Result, issuesOnly bool) (datas []types.Result) {
	var squeezed []types.Result
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
		onlyIssues := make([]types.Result, 0)
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
		var fullFiltering []types.Result
		var fullTreat []types.Result
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
func CheckDoubloons(arr []types.Result) bool {
	for _, a := range arr {
		_, count := Count(arr, a.Address)
		if count != 1 {
			return true
		}
	}

	return false
}

// checks the number of occurrences of the link (Result.Address value)
func Count(res []types.Result, search string) (bool, int) {
	r := len(res)
	count := 0
	for i := 0; i < r; i++ {
		if res[i].Address == search {
			count++
		}
	}

	return count == 1, count
}
