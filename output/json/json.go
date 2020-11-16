package json

import (
	c "../../process/types"
	"errors"
	encode "encoding/json"
)

var ErrJsonifying = errors.New("Could not jsonify !")

// Jsonify converts the array of Result into JSON string
func Jsonify(data []c.Result) (s string, err error) {
	bytes, err := encode.MarshalIndent(data, "", "  ")
	if err != nil {
		return string(bytes), ErrJsonifying
	}

	s = string(bytes)

	return s, err
}
