package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := setupRouter()

	startApp(e)
}

func setupRouter() *echo.Echo {
	e := echo.New()

	e.GET("/lb-status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	return e
}

func startApp(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8000"))
}
