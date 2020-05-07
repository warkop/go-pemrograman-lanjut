package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func main()  {
	e := echo.New()

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")


	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("starting", viper.GetString("appName"))
	fmt.Println("waktu:", viper.GetTime("waktu"))

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})

	e.GET("/index", func(context echo.Context) error {
		return context.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":"+viper.GetString("port")))
}
