package events

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/labstack/echo/v4"
)

// Detail ... Get event detail
// @Summary Get event detail
// @Description Get event detail
// @Tags Events
// @param id path string true "Event ID"
// @Success 200 {object} events.EventDetail
// @Failure 404,400,500 {object} responses.MyError
// @Router /events/{id} [get]
func Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	event, err := events.Detail(int64(id))

	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("Invalid Id").Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}
