package main

import (
	"github.com/adisurya/friendly-garbanzo/helpers"
	"github.com/adisurya/friendly-garbanzo/request_handlers/events"
	"github.com/adisurya/friendly-garbanzo/request_handlers/tickets"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/adisurya/friendly-garbanzo/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Ticket booking API documentation
// @version 0.1.0

// @host localhost:11300
// @BasePath /
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
	ticketsRoute.GET("/:id", tickets.Detail)
	ticketsRoute.POST("/booking", tickets.Book)
	ticketsRoute.GET("/inquiry/:id", tickets.Inquiry)
	ticketsRoute.POST("/payment", tickets.Payment)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":11300"))
}
