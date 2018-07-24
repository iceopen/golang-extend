package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!")
}

// 启动方法
func main() {
	// Echo instance
	e := echo.New()
	e.Debug = false
	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
