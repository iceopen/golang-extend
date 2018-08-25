package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/url"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	url1, err := url.Parse("http://172.16.50.134:9200/")
	if err != nil {
		e.Logger.Fatal(err)
	}
	url2, err := url.Parse("http://172.16.50.134:9200/")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			Name: "target 1",
			URL:  url1,
		},
		{
			Name: "target 2",
			URL:  url2,
		},
	}
	rb := middleware.NewRoundRobinBalancer(nil)
	for _, target := range targets {
		rb.AddTarget(target)
	}
	e.Use(middleware.Proxy(rb))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
