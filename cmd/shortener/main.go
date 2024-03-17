package main

import (
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	customMiddleware "github.com/bobopylabepolhk/ypshortener/pkg/middleware"
)

func run() {
	e := echo.New()

	// logger
	l := logger.New()
	e.Use(customMiddleware.LoggerMiddleware(l))

	// gzip
	e.Use(customMiddleware.GzipMiddleware())
	e.Use(defaultMiddleware.Decompress())

	// routers
	shortener.NewRouter(e)

	e.Logger.Fatal(e.Start(config.Cfg.APIURL))
}

func main() {
	config.InitConfig()
	run()
}
