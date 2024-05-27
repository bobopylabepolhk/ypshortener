package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/db"
	"github.com/bobopylabepolhk/ypshortener/internal/app/healthcheck"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	customMiddleware "github.com/bobopylabepolhk/ypshortener/pkg/middleware"
)

func run() {
	e := echo.New()

	// middleware
	e.Use(customMiddleware.AuthMiddleware(config.Cfg.Secret))

	l := logger.New()
	e.Use(customMiddleware.LoggerMiddleware(l))

	e.Use(customMiddleware.GzipMiddleware())
	e.Use(defaultMiddleware.Decompress())

	// db
	postgres, err := db.New()
	if err != nil {
		logger.Error("failed to init db")
	}

	err = db.Migrate(postgres)
	if err != nil {
		logger.Error("failed to run migrations")
	}

	// routers
	shortener.NewRouter(e, postgres)
	healthcheck.NewRouter(e, postgres)

	// start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		e.Logger.Fatal(e.Start(config.Cfg.APIURL))
	}()

	// shutdown on interrupt
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		err = fmt.Errorf("shutdown failed: %w", err)
		e.Logger.Fatal(err)
	}
}

func main() {
	config.InitConfig()
	run()
}
