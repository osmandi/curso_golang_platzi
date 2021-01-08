package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// Instanciar echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
