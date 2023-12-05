package main

import (
	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitFromCLI()
	e := echo.New()
	shortener.NewRouter(e)

	e.Logger.Fatal(e.Start(config.APIURL))
}
