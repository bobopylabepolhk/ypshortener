package jsonreader

import (
	"errors"
	"os"
)

type JSONReader struct {
	file *os.File
}

var ErrFailedToOpenFile = errors.New("failed to open json file")

func NewJSONReader(path string) (*JSONReader, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, ErrFailedToOpenFile
	}

	return &JSONReader{file: file}, nil
}

func (jr JSONReader) WriteRow(data interface{}) error {
	return nil
}

func (jr JSONReader) FindByKey(key string) ([]byte, error) {
	return make([]byte, 0), nil
}
