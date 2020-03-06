package main

import (
	"flag"
	"github.com/YonghoChoi/slack-interact-api/cmd/api/conf"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./conf/config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is slack interactive api server")
	})

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.0.0")
	})

	e.GET("/slack/approve", Approve)

	e.Logger.Fatal(e.Start(":8000"))
}
