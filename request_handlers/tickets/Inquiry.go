package tickets

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/adisurya/friendly-garbanzo/database/bookings"
	"github.com/labstack/echo/v4"
)

// Inquiry ... Inquiry booking data
// @Summary Inquiry booking data
// @Description Inquiry booking data
// @Tags Tickets
// @param id path string true "Booking ID"
// @Success 200 {object} bookings.BookingInquiry
// @Failure 404,400,500 {object} responses.MyError
// @Router /tickets/inquiry/{id} [get]
func Inquiry(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	inquiry, err := bookings.Inquiry(int64(id))

	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("Unpaid booking not found").Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, inquiry)
}
