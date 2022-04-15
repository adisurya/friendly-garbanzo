package main

import (
	"github.com/adisurya/friendly-garbanzo/request_handlers/events"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	eventRoute := e.Group("/events")
	eventRoute.GET("", events.Index)
	eventRoute.POST("/create", events.Create)

	e.Logger.Fatal(e.Start(":11300"))
}
