package json

import (
	encode "encoding/json"
	c "../../collect"
	"errors"
)

type JsonResult struct {
	index   int
	content c.Result
}

var ErrJsonifying = errors.New("Could not jsonify !")

func Jsonify(data []c.Result) (string, error) {
	formattedData := JsonResult{}
	for i := 0; i < len(data); i++ {
		formattedData.index = i
		formattedData.content = data[i]
	}

	bytes, err := encode.Marshal(formattedData)
	if err != nil {
		return string(bytes), ErrJsonifying
	}

	j := string(bytes)

	return j, err
}