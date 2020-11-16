package json

import (
	c "../../process/types"
	errors "golang.org/x/xerrors"
	"strings"
	"testing"
)

func TestJsonify(t *testing.T) {
	t.Run("Jsonify jsonifies one line !!!", func(t *testing.T) {
		datas := []c.Result{
			{
				Address: "https://www.iana.org/domains/example",
				Status:  "200 OK",
				Code:    200,
				Source:  "http://example.com",
			},
		}
		got, _ := Jsonify(datas)
		want := "{\n  Address: \"https://www.iana.org/domains/example\",\n  Status:  \"200 OK\",\n  Code:    200,\n  Source:  \"http://example.com\",\n}"
		if strings.Compare(got, want) != 0 {
			errors.Errorf("Got %s want %s", got, want)
		}
	})

	t.Run("Jsonify jsonifies multiple lines !!!", func(t *testing.T) {
		datas := []c.Result{
			{
				Address: "https://www.iana.org/domains/example",
				Status:  "200 OK",
				Code:    200,
				Source:  "http://example.com",
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
		got, _ := Jsonify(datas)
		want := ""
		if got != want {
			errors.Errorf("")
		}
	})
}
