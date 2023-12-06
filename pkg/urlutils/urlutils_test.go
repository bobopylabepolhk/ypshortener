package urlutils_test

import (
	"testing"

	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
	"github.com/stretchr/testify/require"
)

func TestValidateUrl(t *testing.T) {
	t.Run("returns true for valid urls", func(t *testing.T) {
		urls := []string{
			"https://browser.yandex.ru/corp/",
			"https://practicum.yandex.ru/trainer/go-advanced/lesson/adfc3335-796e-4df6-b454-6c602a749003/",
			"http://service.app.io?search=a&id=12",
			"https://companies.com?page=0&sort_by=&sort_order=asc&filters=%5B%7B%22column_name%22%3A%22sector%22%2C%22operator%22%3A%22in%22%2C%22value%22%3A%5B%22Cruise%20lines%22%5D%7D%2C%7B%22column_name%22%3A%22subscription_expires_at%22%2C%22operator%22%3A%22in%22%2C%22value%22%3A%5B%222020-06-19%22%2C%222035-06-19%22%5D%7D%5D&page_size=100&search=test",
		}

		for _, v := range urls {
			r := urlutils.ValidateURL(v)
			require.True(t, r, v)
		}
	})

	t.Run("returns false for invalid urls; IPV4/6 addresses; localhost", func(t *testing.T) {
		urls := []string{
			"htps://google.com",
			"http:/vk.ru/",
			"http://service.app.iosearch&id=12",
			"http://4.123.246.235",
			"http://localhost:3333",
		}

		for _, v := range urls {
			r := urlutils.ValidateURL(v)
			require.False(t, r, v)
		}
	})
}
