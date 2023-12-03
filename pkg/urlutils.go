package urlutils

import (
	"regexp"
)

const ValidUrlSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ValidateUrl(url string) bool {
	r := regexp.MustCompile(`^(https?://)?([\da-z.-]+)\.([a-z.]{2,6})([/\w.-]*)*/?$`)

	return r.MatchString(url)
}

func ValidatePathParam(url string) bool {
	r := regexp.MustCompile(`^/[A-Za-z0-9]+/?$`)

	return r.MatchString(url)
}
