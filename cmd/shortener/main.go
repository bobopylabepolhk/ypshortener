package main

import (
	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/pkg/echoeasyjson"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	"github.com/bobopylabepolhk/ypshortener/pkg/middleware"
)

func run() {
	// setup echo
	e := echo.New()
	e.JSONSerializer = echoeasyjson.EasyJSONSerializer{}

	// logger
	l := logger.New()
	e.Use(middleware.LoggerMiddleware(l))

	// routers
	shortener.NewRouter(e)

	e.Logger.Fatal(e.Start(config.Cfg.APIURL))
}

func main() {
	config.InitConfig()
	run()
}
