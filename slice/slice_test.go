package slice

import (
	"testing"
)

func TestSlice(t *testing.T) {

	t.Run("Valid string with one link gives one link", func(t *testing.T) {

		got, _ := Slice("<a href=\"https://example.com\"></a>")
		want := "https://example.com"
		if got[0] != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Valid string with two links gives two links", func(t *testing.T) {
		link, _ := Slice("<a href=\"https://example.com\"></a><a href=\"https://perdu.com\"></a>")
		got := link[0]
		want := "https://example.com"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		gotTwo := link[1]
		wantTwo := "https://perdu.com"
		if gotTwo != wantTwo {
			t.Errorf("got %s want %s", gotTwo, wantTwo)
		}
	})

}