package main

import (
	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/controller"
	_ "github.com/betorvs/sensubot/gateway/sensuclient"
	_ "github.com/betorvs/sensubot/gateway/slackclient"
	_ "github.com/betorvs/sensubot/gateway/telegramclient"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	g := e.Group("/sensubot/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)
	g.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	g.POST("/slack", controller.SlackEvents)
	g.POST("/telegram", controller.TelegramEvents)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
