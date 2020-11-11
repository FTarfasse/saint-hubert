package cli

import (
	c "../collect"
	"reflect"
	"strings"
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

//func TestPrintLine(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		datas := []c.Result{
//			{
//				Address: "abcde",
//				Status:  "200 OK",
//				Code:    200,
//				Source:  "example.com",
//			},
//			{
//				Address: "a",
//				Status:  "a",
//				Code:    1,
//				Source:  "Kittycat",
//			},
//			{
//				Address: "a2",
//				Status:  "abc",
//				Code:    12,
//				Source:  "Kit",
//			},
//		}
//		sizes := []int{5, 6, 3, 11}
//	})
//}

func TestColorOutput(t *testing.T) {
	data := []c.Result{
		{
			Address: "abcde",
			Status:  "200 OK",
			Code:    200,
			Source:  "example.com",
		},
		{
			Address: "a",
			Status:  "404 NOT FOUND",
			Code:    404,
			Source:  "Kittycat",
		},
	}
	got := ColorOutput(data)
	//want := []c.Result{
	//	{
	//		Address: "abcde",
	//		Status:  "\033[0;32m200 OK\033[0m",
	//		Code:    200,
	//		Source:  "example.com",
	//	},
	//	{
	//		Address: "a",
	//		Status:  "\033[1;31m404 NOT FOUND\033[0m",
	//		Code:    404,
	//		Source:  "Kittycat",
	//	},
	//}
	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("\ngot  %v \nwant %v", got, want)
	//}
	if strings.Contains(got[0].Status, "\033[0;32m200 OK\033[0m") == false {
		t.Error("Not properly formatted")
	}
	if strings.Contains(got[1].Status, "\033[1;31m404 NOT FOUND\033[0m") == false {
		t.Error("Not properly formatted")
	}
}