package events

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/responses"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	events, err := events.List()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := responses.EventList{}
	res.Events = events
	return c.JSON(http.StatusOK, res)
}
