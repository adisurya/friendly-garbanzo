package events

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/responses"
	"github.com/labstack/echo/v4"
)

// Index ... Get all events
// @Summary Get all events
// @Description Get all events
// @Tags Events
// @Success 200 {object} responses.EventList
// @Failure 404,400,500 {object} responses.MyError
// @Router /events [get]
func Index(c echo.Context) error {
	events, err := events.List()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := responses.EventList{}
	res.Events = events
	return c.JSON(http.StatusOK, res)
}
