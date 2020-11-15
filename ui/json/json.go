package json

import (
	c "../../collect"
	"errors"
	encode "encoding/json"
	//"fmt"
)

type JsonResult struct {
	index   int
	content c.Result
}

var ErrJsonifying = errors.New("Could not jsonify !")

func Jsonify(data []c.Result) (s string, err error) {
	bytes, err := encode.MarshalIndent(data, "", "  ")
	if err != nil {
		return string(bytes), ErrJsonifying
	}

	s = string(bytes)

	return s, err
}
