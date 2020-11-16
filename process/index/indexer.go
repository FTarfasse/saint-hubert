// slice package contains the helper function to slice the links from the HTML string
package slice

import (
	"strings"
)

const (
	href    string = "href"
	http    string = "http"
	endLink string = "\""
)

var tmp string

// utility function to target the links into the HTML string
// It targets the href elements having http or https links
func Indexer(s string) (start, end int) {
	iHref := strings.Index(s, href)
	tmp = s[iHref:len(s)]
	iHttp := strings.Index(tmp, http)
	tmp := s[iHref+iHttp : len(s)]
	endUrlIndex := strings.Index(tmp, endLink)
	start = iHttp + iHref
	end = endUrlIndex + start

	return
}