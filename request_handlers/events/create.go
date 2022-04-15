package events

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
	"github.com/adisurya/friendly-garbanzo/structs/responses"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	e := new(requests.CreateEvent)
	if err := c.Bind(e); err != nil {
		return err
	}
	println("request_handlers/events/create.go:Create(): " + e.EventTime)
	id, err := events.Create(e)
	if err != nil {
		return err
	}

	res := new(responses.ResponseId)
	res.Id = id
	return c.JSON(http.StatusCreated, res)
}
