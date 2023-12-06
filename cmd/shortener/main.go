package main

import (
	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
)

func main() {
	config.InitConfig()
	e := echo.New()
	shortener.NewRouter(e)

	e.Logger.Fatal(e.Start(config.Cfg.APIURL))
}
