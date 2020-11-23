package cli

import (
	"../../process/types"
	"github.com/acarl005/stripansi"
	"reflect"
	"strings"
	"testing"
)

func TestColumnsSizes(t *testing.T) {
	t.Run("Columns should give the length of each struct field", func(t *testing.T) {
		datas := []types.Result{
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

func TestBuildTable(t *testing.T) {
	t.Run("\nTable build properly", func(t *testing.T) {
		datas := []types.Result{
			{
				Address: "abc",
				Status:  "OK",
				Code:    200,
				Source:  "example.com",
			},
		}
		maxs := []int{10, 10}
		got := BuildTable(ColorOutput(datas), maxs)
		want := []string{
			"| Link       | Status     |\n",
			"| abc        | \u001B[0;32mOK\u001B[0m         |\n",
		}
		for i := 0; i < 2; i++ {
			got[i] = stripansi.Strip(got[i])
			want[i] = stripansi.Strip(want[i])
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nGOT\n  %v \nWANT\n %v", got, want)
		}
	})
}

func TestBuildHeader(t *testing.T) {
	t.Run("\nBuild line properly", func(t *testing.T) {

		msg := []string{"Link", "Status"}
		sizes := []int{
			10,
			10,
		}
		got := BuildHeader(msg, Separator, sizes)
		want := "| Link       | Status     |"
		if got != want {
			t.Errorf("\n Got:  %s\n Want: %s\n", got, want)
		}
	})
}

func TestBuildLine(t *testing.T) {
	t.Run("\nBuild line properly", func(t *testing.T) {

		msg := []string{"abc", "\u001B[0;32mOK\u001B[0m"}
		//fmt.Printf("\nREAL LENGTH: %v\n", len(msg[1]))
		sizes := []int{
			10,
			10,
		}

		got := BuildLine(msg, Separator, sizes)
		want := "| abc        | \u001B[0;32mOK\u001B[0m         |"
		if got != want {
			t.Errorf("\n Got:  %s\n Want: %s\n", got, want)
		}
	})
}

func TestColorOutput(t *testing.T) {
	data := []types.Result{
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

	if len(got[0].Status) != len("\033[0;32m200 OK\033[0m") {
		t.Error("Not properly formatted")
	}
	if !strings.Contains(got[0].Status, "200 OK") {
		t.Error("OK 200 content corrupted")
	}

	if len(got[1].Status) != len("\033[1;31m404 NOT FOUND\033[0m") { // 30 - 6
		//if strings.Contains(got[1].Status, "\033[1;31m404 NOT FOUND\033[0m") == false {
		t.Error("Not properly formatted")
	}

	if !strings.Contains(got[1].Status, "404 NOT FOUND") {
		t.Error("404 NOT FOUND content corrupted")
	}

}
