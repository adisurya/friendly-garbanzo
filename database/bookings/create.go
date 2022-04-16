package bookings

import (
	"strconv"
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
)

type CreateBookingResponse struct {
	Id           int64 `json:"id"`
	TotalTicket  int   `json:"total_ticket"`
	TotalPayment int   `json:"total_payment"`
}

func Create(data *requests.CreateBooking) (CreateBookingResponse, error) {
	res := CreateBookingResponse{}

	db, err := database.Connect()
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}
	defer db.Close()

	event, err := events.Detail(data.EventId)
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}

	stmtInst, err := db.Prepare(`
		INSERT INTO bookings(
			id, name, event_id, total_booked, ticket_price, total_price,
			valid_until, is_paid, created_at)
			VALUES(NULL, ?, ?, ?, ?, ?, ?, 0, now())
			`)
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}
	defer stmtInst.Close()

	ticketPrice, err := strconv.Atoi(event.TicketPrice)
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}

	now := time.Now().Add(time.Hour * 5)

	totalPrice := data.Total * ticketPrice
	result, err := stmtInst.Exec(data.Name, data.EventId, data.Total, event.TicketPrice, totalPrice, now)
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return res, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return res, err
	}

	res.Id = insertId
	res.TotalTicket = data.Total
	res.TotalPayment = totalPrice

	return res, nil
}
