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
