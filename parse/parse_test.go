package parse

import (
	"testing"
)

func TestParse(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Valid string with no link gives alert message", func(t *testing.T) {
		got := Parse("<body><a>YOLO</a></body>")
		want := "No links at this address !"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Valid string with one link gives one link", func(t *testing.T) {
		got := Parse("aefaefef<a href=\"https://example.com\"></a>")
		want := "https://perdu.com"
		assertCorrectMessage(t, got, want)
	})

}
