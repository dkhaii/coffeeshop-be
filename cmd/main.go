package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {	
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Selamat datang di echo!")
	})

	app.Logger.Fatal(app.Start(":8080"))
}
