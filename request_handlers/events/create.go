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

	if err := c.Validate(e); err != nil {
		return err
	}

	id, err := events.Create(e)
	println("request_handlers/events/create.go:Create(): " + err.Error())

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := new(responses.ResponseId)
	res.Id = id
	return c.JSON(http.StatusCreated, res)
}
