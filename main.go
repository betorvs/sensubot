package main

import (
	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/controller"
	_ "github.com/betorvs/sensubot/gateway/customlog"
	_ "github.com/betorvs/sensubot/gateway/sensuclient"
	_ "github.com/betorvs/sensubot/gateway/slackclient"
	_ "github.com/betorvs/sensubot/gateway/telegramclient"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
