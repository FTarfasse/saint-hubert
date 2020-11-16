package slice

import (
	index "../index"
	valid "../validate"
	"errors"
)

var ErrNoLinksFound = errors.New("No links in this source !")

func SliceLinks(s string) (links []string, err error) {
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