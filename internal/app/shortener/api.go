package shortener

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bobopylabepolhk/ypshortener/config"
	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

func handleGetURL(us *URLShortener, w http.ResponseWriter, r *http.Request) {
	if !urlutils.ValidatePathParam(r.URL.Path) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := strings.Replace(r.URL.Path, "/", "", 1)

	ogURL, err := us.GetOriginalURL(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Location", ogURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleShortenURL(us *URLShortener, w http.ResponseWriter, r *http.Request) {
	ogURL, err := io.ReadAll(r.Body)
	if r.URL.Path != "/" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := us.GetShortURLToken()
	err = us.SaveShortURL(string(ogURL), token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("%s/%s", config.APIURL, token)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func handleShortener(us *URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			{
				handleGetURL(us, w, r)
			}
		case http.MethodPost:
			{
				handleShortenURL(us, w, r)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func Router(m *http.ServeMux) {
	us := NewURLShortener(6)
	m.HandleFunc("/", handleShortener(us))
}
