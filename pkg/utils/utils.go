package utils

import (
	"encoding/json"
	"io"
)

func ParseBody(jsonData io.Reader, output interface{}) error {

	decoder := json.NewDecoder(jsonData)
	err := decoder.Decode(output)

	if err != nil {
		return err
	}

	return nil
}
