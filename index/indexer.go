package slice

import (
	"strings"
)

var tmp string

func Indexer(s, sub string) (start, end int) {
	href := strings.Index(s, sub)
	//fmt.Sprintf("href %v \n", href)
	tmp = s[href:len(s)]
	//fmt.Printf("tmp %s \n", tmp)
	http := strings.Index(tmp, "http")
	//fmt.Printf("http %v\n", http)
	tmp := s[href+http:len(s)]
	//fmt.Printf("new tmp %v \n", tmp)
	endUrlIndex := strings.Index(tmp, "\"")
	//fmt.Printf("tmp[:endUrlIndex] %v\n", tmp[:endUrlIndex])
	start = http + href
	//fmt.Printf("start %v\n", start)
	end = endUrlIndex + http + href
	//fmt.Printf("end %v\n", end)
	//fmt.Printf("FINAL TEST %v\n", s[start:end])
	//start = strings.Index(s[start:len(s)-start], "http")
	// fmt.Println(s[start:end])
	// end = strings.Index(tmp, "http")
	// tmpSecond := s[end:len(s) - start]
	// fmt.Println(tmpSecond)
	return
}
