package slice

import (
	"testing"
)

func TestIndexer(t *testing.T) {
	t.Run("Indexer retrieve starting index of href=\"", func(t *testing.T) {
		start, _ := Indexer("<a href=\"https://example.com\"></a>", "href=")
		got := start
		want := 9
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Indexer retrieve ending index of href=\"", func(t *testing.T) {
		_, end := Indexer("<a href=\"https://example.com\"></a>", "href=")
		got := end
		want := 28
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Indexer retrieve indexes of a link", func(t *testing.T) {
		s := "<a href=\"https://example.com\"></a>"
		start, end := Indexer(s, "href=")
		got := s[start:end]
		want := "https://example.com"
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
