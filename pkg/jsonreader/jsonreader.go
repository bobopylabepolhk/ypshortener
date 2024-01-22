package jsonreader

import (
	"encoding/json"
	"errors"
	"os"
)

type JSONReader struct {
	file    *os.File
	decoder *json.Decoder
}

var ErrFailedToOpenFile = errors.New("failed to open json file")

func NewJSONReader(path string) (*JSONReader, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		return nil, ErrFailedToOpenFile
	}

	return &JSONReader{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (jr JSONReader) InitFromFile() ([]map[string]interface{}, error) {
	data := []map[string]interface{}{}
	for jr.decoder.More() {
		row := map[string]interface{}{}
		err := jr.decoder.Decode(&row)
		if err != nil {
			return data, err
		}

		data = append(data, row)
	}

	return data, nil
}

func (jr JSONReader) WriteRow(row interface{}) error {
	data, err := json.Marshal(&row)
	if err != nil {
		return err
	}

	_, err = jr.file.Write(append(data, '\n'))
	return err
}
