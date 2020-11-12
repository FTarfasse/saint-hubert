package json

import (
	errors "golang.org/x/xerrors"
	"testing"
	c "../../collect"
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
			//{
			//	Address: "a",
			//	Status:  "a",
			//	Code:    1,
			//	Source:  "Kittycat",
			//},
			//{
			//	Address: "a2",
			//	Status:  "abc",
			//	Code:    12,
			//	Source:  "Kit",
			//},
		}
		b, _ := Jsonify(datas)
		got := string(b)
		want := "{index: 0, content: {\n\t\t\t\tAddress: \"https://www.iana.org/domains/example\",\n\t\t\t\tStatus:  \"200 OK\",\n\t\t\t\tCode:    200,\n\t\t\t\tSource:  \"http://example.com\",\n\t\t\t}}"
		if got != want {
			errors.Errorf("")
		}
	})
}
