package cli

import (
	"fmt"
	"testing"
	c "../collect"
)

func TestColumns(t *testing.T) {
	t.Run("Columns should give the length of each struct field", func(t *testing.T) {
		got := Columns(c.Result{
			Address: "abcde",
			Status:  "200 OK",
		})
		want := []int{5, 6}
		fmt.Println(want)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

//func TestMaxer(t *testing.T) {
//
//	t.Run("Maxer should find the max length of each struct field", func(t *testing.T) {
//		sizes := Columns()
//		got :=
//		want := []int{2, 4}
//		if got != want {
//			t.Errorf("got %v want %v", got, want)
//		}
//})
//}
