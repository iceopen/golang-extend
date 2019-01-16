package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e = echo.New()

func init() {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	health()
}

func Start(address string) {
	e.Logger.Fatal(e.Start(address))
}

func health() {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "UP")
	})
}
