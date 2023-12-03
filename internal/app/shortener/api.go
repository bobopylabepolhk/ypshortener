package shortener

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bobopylabepolhk/ypshortener/config"
	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

func handleGetURL(w http.ResponseWriter, r *http.Request) {
	if !urlutils.ValidatePathParam(r.URL.Path) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := strings.Replace(r.URL.Path, "/", "", 1)

	ogURL, err := GetOriginalURL(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Location", ogURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	ogURL, err := io.ReadAll(r.Body)
	if r.URL.Path != "/" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := GetShortURLToken(6)
	err = SaveShortURL(string(ogURL), token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("%s/%s", config.APIURL, token)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func handleShortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			handleGetURL(w, r)
		}
	case http.MethodPost:
		{
			handleShortenURL(w, r)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Router(m *http.ServeMux) {
	m.HandleFunc("/", handleShortener)
}
