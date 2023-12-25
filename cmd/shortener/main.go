package main

import (
	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	"github.com/bobopylabepolhk/ypshortener/pkg/middleware"
)

func run() {
	e := echo.New()
	l, err := logger.New()

	if err != nil {
		e.Logger.Error("couldn't initialize zap logger")
	}

	e.Use(middleware.LoggerMiddleware(l))

	shortener.NewRouter(e)

	e.Logger.Fatal(e.Start(config.Cfg.APIURL))
}

func main() {
	config.InitConfig()
	run()
}
