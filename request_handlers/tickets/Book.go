package tickets

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/bookings"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
	"github.com/labstack/echo/v4"
)

// Book ... Booking tickets for an event
// @Summary Booking tickets for an event
// @Description Booking tickets for an event
// @Tags Tickets
// @param event body requests.CreateBooking true "Booking Data"
// @Success 201 {object} bookings.CreateBookingResponse
// @Failure 404,400,500 {object} responses.MyError
// @Router /tickets/booking [post]
func Book(c echo.Context) error {
	b := new(requests.CreateBooking)
	if err := c.Bind(b); err != nil {
		return err
	}

	if err := c.Validate(b); err != nil {
		return err
	}

	book, err := bookings.Create(b)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}
