package main

import (
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	shortener.NewRouter(e)

	addr := fmt.Sprintf(":%d", config.PORT)
	e.Logger.Fatal(e.Start(addr))
}
