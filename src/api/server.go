package api

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/lmittmann/tint"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func MainServer(port string) {
	e := echo.New()
	prettyHandler := tint.NewTextHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.TimeOnly,
	})
	e.Logger = slog.New(prettyHandler)
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	SetUp(e)
	e.GET("/health", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "server healthy"})
	})
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start simulator server")
	}
	e.Logger.Info("Simulator server started successfully")

}
