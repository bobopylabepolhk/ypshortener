package urlutils

import (
	"regexp"
)

func ValidateURL(url string) bool {
	r := regexp.MustCompile(`^(https?://)?([\da-z.-]+)\.([a-z.]{2,6})([/\w\.-]*)*(\?[^\s#]*)?(#[^\s]*)?$`)

	return r.MatchString(url)
}

// TODO remove func after rewriting to echo

func ValidatePathParam(url string) bool {
	r := regexp.MustCompile(`^/[A-Za-z0-9]+/?$`)

	return r.MatchString(url)
}
