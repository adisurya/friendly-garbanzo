package events

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
	"github.com/adisurya/friendly-garbanzo/structs/responses"
	"github.com/labstack/echo/v4"
)

// Create ... Create event
// @Summary Create event
// @Description Create event
// @Tags Events
// @param event body requests.CreateEvent true "Event Data"
// @Success 201 {object} responses.ResponseId
// @Failure 404,400,500 {object} responses.MyError
// @Router /events [post]
func Create(c echo.Context) error {
	e := new(requests.CreateEvent)
	if err := c.Bind(e); err != nil {
		return err
	}

	if err := c.Validate(e); err != nil {
		return err
	}

	id, err := events.Create(e)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := new(responses.ResponseId)
	res.Id = id
	return c.JSON(http.StatusCreated, res)
}
