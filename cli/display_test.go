package cli

import (
	c "../collect"
	"reflect"
	"testing"
)

func TestColumnsSizes(t *testing.T) {
	t.Run("Columns should give the length of each struct field", func(t *testing.T) {
		datas := []c.Result{
			{
				Address: "abcde",
				Status:  "200 OK",
				Code:    200,
				Source:  "example.com",
			},
			{
				Address: "a",
				Status:  "a",
				Code:    1,
				Source:  "Kittycat",
			},
			{
				Address: "a2",
				Status:  "abc",
				Code:    12,
				Source:  "Kit",
			},
		}
		got := ColumnsSizes(datas)
		want := []int{5, 6, 3, 11}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

//func TestMaxElemLengthInArray(t *testing.T) {
//	t.Run("Maxer should find the max index for above the arrays", func(t *testing.T) {
//		got := Max([]Result)
//		want := []int{7, 8, 4, 12}
//		if got != want {
//			t.Errorf("got %v want %v", got, want)
//		}
//	})
//}

//func TestPrintTableLine(t *testing.T) {
//	t.Run("CliFormat should prepare header column size ", func(t *testing.T) {
//		link := "Link"
//		status := "Status"
//		sizes := CliFormat(4, 5, link, status)
//		got := " _________\n|" + link + "" | ____ | ____ |\n
//		"
//		want := []int{2, 4}
//		if got != want {
//			t.Errorf("got %v want %v", got, want)
//		}
//	})
//}
