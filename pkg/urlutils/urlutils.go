package urlutils

import (
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func ValidateURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	// for invalid params without ?
	if u.RawQuery == "" && strings.Contains(rawURL, "&") {
		return false
	}

	host := u.Hostname()
	validScheme := u.Scheme == "http" || u.Scheme == "https"
	validHost := strings.Contains(host, ".")
	isLocal := host == "127.0.0.1" || host == "0.0.0.0"

	return validScheme && validHost && !isLocal
}

func CreateRandomToken(tokenLen int) string {
	const tokenChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	result := make([]byte, tokenLen)

	for i := 0; i < tokenLen; i++ {
		result[i] = tokenChars[rand.Intn(len(tokenChars))]
	}

	return string(result)
}
