package router

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-redis-event-listener/api/httpserver/router/healthcheckrouter"
	"github.com/saeedjhn/go-redis-event-listener/api/httpserver/router/userrouter"
	"github.com/saeedjhn/go-redis-event-listener/internal/bootstrap"
)

func Setup(
	app *bootstrap.Application,
	echo *echo.Echo,
) {
	routerGroup := echo.Group("")

	userrouter.New(app, routerGroup)
	healthcheckrouter.New(app, routerGroup)
}
