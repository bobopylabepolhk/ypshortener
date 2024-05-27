package repo

import (
	"errors"
	"fmt"
)

var ErrDuplicateURL = errors.New("shortURL already exists for this ogURL")
var ErrURLIsDeleted = errors.New("shortURL is deleted")

func errShortURLDoesNotExist(shortURL string) error {
	return fmt.Errorf("short url %s was never created", shortURL)
}

func errOgURLNotFound(ogURL string) error {
	return fmt.Errorf("original url %s was not found", ogURL)
}
