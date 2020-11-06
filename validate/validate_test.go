package validate

import "testing"

func TestValidator(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("Return false if string is empty", func(t *testing.T) {
		got := Validate("")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return true if it contains http://", func(t *testing.T) {
		got := Validate("href=\"http://")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return false if it does not contain http", func(t *testing.T) {
		got := Validate("let's go bananas with this test string !!!")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return true if it contains https://", func(t *testing.T) {
		got := Validate("href=\"https://")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Return false if it does not contain https://", func(t *testing.T) {
		got := Validate("let's go bananas with this test string !!!")
		want := false
		assertCorrectMessage(t, got, want)
	})
}