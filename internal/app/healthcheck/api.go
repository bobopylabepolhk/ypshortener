package healthcheck

import (
	"database/sql"
	"net/http"

	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	"github.com/labstack/echo/v4"
)

type (
	Healthcheck interface {
		PingDB() error
	}

	Router struct {
		HealthcheckService Healthcheck
	}
)

func (router *Router) HandlePing(ctx echo.Context) error {
	err := router.HealthcheckService.PingDB()

	if err != nil {
		logger.Error(err.Error())
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusOK)
}

func NewRouter(e *echo.Echo, db *sql.DB) {
	repo := NewHealthcheckRepo(db)
	hs := NewHealthcheckService(repo)
	router := &Router{HealthcheckService: hs}

	e.GET("/ping", router.HandlePing)

}
