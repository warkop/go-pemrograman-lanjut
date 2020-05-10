package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/secure"
)

func main() {
	e := echo.New()

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:7000", "www.google.com"},
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))
	e.GET("/index", func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.StartTLS(":7000", "server.crt", "server.key"))
}
