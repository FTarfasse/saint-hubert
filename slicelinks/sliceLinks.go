package slice

import (
	index "../index"
	valid "../validate"
	"errors"
)

var ErrNoLinksFound = errors.New("No links in this source !")

func SliceLinks(s string) (links []string, err error) {

	// for init_part; condition_part; post_part {
	// 	...
	// }
	// The init part is executed first before the first iteration
	// The condition part is executed before every iteration. If the condition is false the loop will exit otherwise the loop will continue to iterate.
	// The post part is executed after every iteration. After this the condition is check, if it is true then the loop is continued otherwise loop exists.

	var sliced string

	for valid.Validate(s) {
		start, end := index.Indexer(s)
		link := s[start:end]
		sliced = s[end:]
		links = append(links, link)
		s = sliced
	}

	if len(links) == 0 {
		return links, ErrNoLinksFound
	}

	return
}