package urlutils

import (
	"math/rand"
	"regexp"
	"time"
)

func ValidateURL(url string) bool {
	r := regexp.MustCompile(`^(https?://)?([\da-z.-]+)\.([a-z.]{2,6})([/\w\.-]*)*(\?[^\s#]*)?(#[^\s]*)?$`)

	return r.MatchString(url)
}

func GetShortURLToken() string {
	const tokenChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	tokenLen := 6
	result := make([]byte, tokenLen)

	for i := 0; i < tokenLen; i++ {
		result[i] = tokenChars[rand.Intn(len(tokenChars))]
	}

	return string(result)
}
