package shortener

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bobopylabepolhk/ypshortener/config"
	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

func handleGetUrl(w http.ResponseWriter, r *http.Request) {
	if !urlutils.ValidatePathParam(r.URL.Path) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := strings.Replace(r.URL.Path, "/", "", 1)

	ogUrl, err := GetOriginalUrl(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Location", ogUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleShortenUrl(w http.ResponseWriter, r *http.Request) {
	ogUrl, err := io.ReadAll(r.Body)
	if r.URL.Path != "/" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := GetShortUrlToken(6)
	err = SaveShortUrl(string(ogUrl), token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("%s/%s", config.API_URL, token)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func handleShortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			handleGetUrl(w, r)
		}
	case http.MethodPost:
		{
			handleShortenUrl(w, r)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Router(m *http.ServeMux) {
	m.HandleFunc("/", handleShortener)
}
