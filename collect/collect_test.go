package collect

import (
	"reflect"
	"testing"
)

func TestCollect(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertEquals := func(t *testing.T, got, want int) {
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	//t.Run("Valid string with no link gives alert message", func(t *testing.T) {
	//	got := Collect("http://perdu.com")
	//	want := "No links at this address !"
	//	assertCorrectMessage(t, got, want)
	//})

	t.Run("Valid string with one link gives one link", func(t *testing.T) {
		got, _ := Collect("https://example.com")
		want := "https://www.iana.org/domains/example"
		assertCorrectMessage(t, got[0].Address, want)
	})

	t.Run("Squeeze erases duplicate links from []Result", func(t *testing.T) {
		ex := []Result{
			{
				Address: "example.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},
			{
				Address: "yolo.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},
			{
				Address: "example.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},
			{
				Address: "yolo.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},
			{
				Address: "yolo.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},

		}
		expected := []Result{
			{
				Address: "example.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},
			{
				Address: "yolo.com",
				Status:  "404 NOT FOUND",
				Code:    404,
				Source:  "none",
			},

		}
		squeezed := Squeeze(ex)
		got := len(squeezed)
		want := len(expected)
		assertEquals(t, got, want)
		if !reflect.DeepEqual(squeezed, expected) {
				t.Errorf("got %v want %v", squeezed, expected)
		}
	})

	t.Run("CheckDoubloons gives false if array has element address more than once", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouf", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "paf", Status: "", Code: 0, Source: ""})
		got := CheckDoubloons(fix)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("CheckDoubloons gives true if array has element address more than once", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouf", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		got := CheckDoubloons(fix)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Count gives true if element is present once", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouf", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "paf", Status: "", Code: 0, Source: ""})
		got, _ := Count(fix, fix[0].Address)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Count gives false if element is present MORE than once", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouf", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		got, _ := Count(fix, fix[0].Address)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Count also gives the number of times the doubloon is present", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouf", Status: "", Code: 0, Source: ""})
		fix = append(fix, Result{Address: "pouet", Status: "", Code: 0, Source: ""})
		_, got := Count(fix, fix[0].Address)
		want := 2
		assertEquals(t, got, want)
	})

}
