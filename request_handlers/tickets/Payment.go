package tickets

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/adisurya/friendly-garbanzo/database/bookings"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
	"github.com/labstack/echo/v4"
)

func Payment(c echo.Context) error {
	p := new(requests.CreatePayment)

	if err := c.Bind(p); err != nil {
		return err
	}

	if err := c.Validate(p); err != nil {
		return err
	}

	inquiryResponse, err := bookings.Inquiry(p.BookingId)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, errors.New("Payment data not found").Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if inquiryResponse.TotalPrice != p.Total {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Payment value does not match").Error())

	}

	now := time.Now()
	if inquiryResponse.ValidUntil.Before(now) {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Expired booking id").Error())

	}

	paymentResponse, err := bookings.Payment(p.BookingId)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, paymentResponse)

}
