package events

import (
	"net/http"
	"strconv"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/labstack/echo/v4"
)

func Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	event, err := events.Detail(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}
