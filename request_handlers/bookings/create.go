package bookings

import (
	"net/http"

	"github.com/adisurya/friendly-garbanzo/database/bookings"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
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
