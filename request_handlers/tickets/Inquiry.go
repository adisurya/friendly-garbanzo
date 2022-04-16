package tickets

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/adisurya/friendly-garbanzo/database/bookings"
	"github.com/labstack/echo/v4"
)

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
