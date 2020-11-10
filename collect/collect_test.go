package collect

import (
	"testing"
)

func TestCollect(t *testing.T) {

	//assertCorrectMessage := func(t *testing.T, got, want string) {
	//	t.Helper()
	//	if got != want {
	//		t.Errorf("got %q want %q", got, want)
	//	}
	//}
	//
	//assertEquals := func(t *testing.T, got, want int){
	//	if got != want {
	//		t.Errorf("got %v want %v", got, want)
	//	}
	//}

	// t.Run("Valid string with no link gives alert message", func(t *testing.T) {
	// 	got := Collect("http://perdu.com")
	// 	fmt.Println(got)
	// 	want := "No links at this address !"
	// 	assertCorrectMessage(t, got, want)
	// })

//	t.Run("Valid string with one link gives one link", func(t *testing.T) {
//		got := Collect("https://example.com")
//		want := "https://www.iana.org/domains/example"
//		assertCorrectMessage(t, got[0].Address, want)
//	})
//
//	t.Run("Erase duplicate links from []Result", func(t *testing.T) {
//		ex := make([]Result, 2)
//		ex = append(ex,Result{
//			Address: "example.com",
//			Status:  "404 NOT FOUND",
//			Code:    404,
//			Source:  "none",
//		})
//		ex = append(ex,Result{
//			Address: "example.com",
//			Status:  "404 NOT FOUND",
//			Code:    404,
//			Source:  "none",
//		})
//		RemoveDuplicates(ex)
//		got := len(ex)
//		want := 1
//		assertEquals(t, got, want)
//	})

	t.Run("Find gives true if element is present", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{ Address: "pouet", Status:  "", Code:    0, Source:  ""})
		fix = append(fix, Result{ Address: "pouf", Status:  "", Code:    0, Source:  ""})
		fix = append(fix, Result{ Address: "pouet", Status:  "", Code:    0, Source:  ""})
		got := Find(fix, fix[0].Address)
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Find gives false if element is absent", func(t *testing.T) {
		var fix []Result
		fix = append(fix, Result{ Address: "pouet", Status:  "", Code:    0, Source:  ""})
		fix = append(fix, Result{ Address: "pouf", Status:  "", Code:    0, Source:  ""})
		fix = append(fix, Result{ Address: "bananas", Status:  "", Code:    0, Source:  ""})
		got := Find(fix, fix[0].Address)
		want := false
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

}
