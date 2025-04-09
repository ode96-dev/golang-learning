package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
Echo
- High performance and robust middleware ecosystem
	- built-in middleware (gzip, logger, CORS, JWT, etc)
	- grouped routes
	- context-aware handlers
	- easy templating and static file serving
*/

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, echoing!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
