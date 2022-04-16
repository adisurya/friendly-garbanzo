package main

import (
	"github.com/adisurya/friendly-garbanzo/helpers"
	"github.com/adisurya/friendly-garbanzo/request_handlers/events"
	"github.com/adisurya/friendly-garbanzo/request_handlers/tickets"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	eventRoute := e.Group("/events")
	eventRoute.GET("", events.Index)
	eventRoute.GET("/:id", events.Detail)
	eventRoute.POST("", events.Create)

	ticketsRoute := e.Group("/tickets")
	// bookingRoute.GET("", events.Index)
	// bookingRoute.GET("/:id", events.Detail)
	ticketsRoute.POST("booking", tickets.Book)

	e.Logger.Fatal(e.Start(":11300"))
}
