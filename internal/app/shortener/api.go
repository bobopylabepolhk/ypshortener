package shortener

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bobopylabepolhk/ypshortener/config"
	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

type (
	URLShortener interface {
		GetShortURLToken() string
		SaveShortURL(url string, token string) error
		GetOriginalURL(shortURL string) (string, error)
	}

	Router struct {
		Us URLShortener
	}
)

func (router *Router) HandleGetURL(w http.ResponseWriter, r *http.Request) {
	if !urlutils.ValidatePathParam(r.URL.Path) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := strings.Replace(r.URL.Path, "/", "", 1)

	ogURL, err := router.Us.GetOriginalURL(token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Location", ogURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (router *Router) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	ogURL, err := io.ReadAll(r.Body)
	if r.URL.Path != "/" || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := router.Us.GetShortURLToken()
	err = router.Us.SaveShortURL(string(ogURL), token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("%s/%s", config.APIURL, token)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(res))
}

func handleShortener(router *Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			{
				router.HandleGetURL(w, r)
			}
		case http.MethodPost:
			{
				router.HandleShortenURL(w, r)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func NewRouter(m *http.ServeMux) {
	us := NewURLShortenerService()
	router := &Router{Us: us}
	m.HandleFunc("/", handleShortener(router))
}
