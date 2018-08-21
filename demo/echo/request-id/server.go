package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
)

// Handler
func hello(c echo.Context) error {
	fmt.Println(c.Response().Header().Get(echo.HeaderXRequestID))
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
