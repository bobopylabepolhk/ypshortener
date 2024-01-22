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

func (jr JSONReader) InitFromFile() []map[string]interface{} {
	data := []map[string]interface{}{}
	for jr.decoder.More() {
		row := map[string]interface{}{}
		jr.decoder.Decode(&row) // TODO HANDLE ERROR

		data = append(data, row)
	}

	return data
}

func (jr JSONReader) WriteRow(row interface{}) error {
	data, err := json.Marshal(&row)
	if err != nil {
		return err
	}

	_, err = jr.file.Write(append(data, '\n'))
	return err
}
