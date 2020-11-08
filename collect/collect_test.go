package collect

import (
	"testing"
)

func TestCollect(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// t.Run("Valid string with no link gives alert message", func(t *testing.T) {
	// 	got := Collect("http://perdu.com")
	// 	fmt.Println(got)
	// 	want := "No links at this address !"
	// 	assertCorrectMessage(t, got, want)
	// })

	t.Run("Valid string with one link gives one link", func(t *testing.T) {
		got := Collect("https://example.com")
		want := "https://www.iana.org/domains/example"
		assertCorrectMessage(t, got[0].Address, want)
	})

}
