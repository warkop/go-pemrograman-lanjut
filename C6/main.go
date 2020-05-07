package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc  {
	return func(context echo.Context) error {
		fmt.Println("from middleware one")
		return next(context)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		fmt.Println("from middleware two")
		return next(context)
	}
}

func middlewareSomething(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("from middleware something")
		next.ServeHTTP(writer, request)
	})
}

func main()  {
	e := echo.New()
	e.Use(middlewareOne)
	e.Use(middlewareTwo)
	e.Use(echo.WrapMiddleware(middlewareSomething))
	e.GET("/index", func(context echo.Context) error {
		fmt.Println("taeeeeee")
		return context.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}
