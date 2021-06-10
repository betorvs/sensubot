package main

import (
	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/controller"
	_ "github.com/betorvs/sensubot/gateway/customlog"

	_ "github.com/betorvs/sensubot/gateway/chat/googlechat"
	_ "github.com/betorvs/sensubot/gateway/chat/slackclient"
	_ "github.com/betorvs/sensubot/gateway/chat/telegramclient"
	_ "github.com/betorvs/sensubot/gateway/integrations/alertmanagerclient"
	_ "github.com/betorvs/sensubot/gateway/integrations/opsgenieclient"
	_ "github.com/betorvs/sensubot/gateway/integrations/sensuclient"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
