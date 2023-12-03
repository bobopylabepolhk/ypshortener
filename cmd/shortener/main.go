package main

import (
	"net/http"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
)

func run() error {
	mux := http.NewServeMux()
	shortener.Router(mux)

	return http.ListenAndServe(":8080", mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
