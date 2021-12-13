package utils

import (
	"bytes"
	"encoding/gob"
)

func Encode(object interface{}) ([]byte, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(object)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decode(encoded []byte, object interface{}) error {
	buffer := bytes.Buffer{}
	buffer.Write(encoded)
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(object)
	return err
}
