package tickets

import (
	"net/http"
	"strconv"

	"github.com/adisurya/friendly-garbanzo/database/tickets"
	"github.com/labstack/echo/v4"
)

// Detail ... Ticket detail
// @Summary Ticket detail
// @Description Ticket detail
// @Tags Tickets
// @param id path string true "Ticket ID"
// @Success 200 {object} tickets.TicketDetailStruct
// @Failure 404,400,500 {object} responses.MyError
// @Router /tickets/{id} [get]
func Detail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	event, err := tickets.Detail(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)

}
