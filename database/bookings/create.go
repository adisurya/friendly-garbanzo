package bookings

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
	"github.com/adisurya/friendly-garbanzo/database/events"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
)

type CreateBookingResponse struct {
	Id           int64     `json:"id"`
	TotalTicket  int       `json:"total_ticket"`
	TotalPayment int       `json:"total_payment"`
	ValidUntil   time.Time `json:"valid_until"`
}

func Create(data *requests.CreateBooking) (CreateBookingResponse, error) {
	res := CreateBookingResponse{}

	db, err := database.Connect()
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}
	defer db.Close()

	trx, err := db.Begin()
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}

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
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		trx.Rollback()

		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}
	err = updateTicket(insertId, data.EventId, data.Total, db)
	if err != nil {
		trx.Rollback()
		println("database/bookings/create.go:Create(): " + err.Error())
		return res, err
	}

	res.Id = insertId
	res.TotalTicket = data.Total
	res.TotalPayment = totalPrice
	res.ValidUntil = now

	err = trx.Commit()
	if err != nil {
		println("database/bookings/create.go:Create(): " + err.Error())

		return res, err
	}

	return res, nil
}

func updateTicket(id int64, eventId int64, totalTickets int, db *sql.DB) error {

	stmtUpd, err := db.Prepare("update tickets set booking_id = ? WHERE event_id = ? AND booking_id IS NULL LIMIT ?")
	if err != nil {
		println("database/events/create.go:updateTicket(): " + err.Error())
		return err
	}
	defer stmtUpd.Close()

	result, err := stmtUpd.Exec(id, eventId, totalTickets)
	if err != nil {
		println("database/events/create.go:updateTicket(): " + err.Error())
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		println("database/events/create.go:updateTicket(): " + err.Error())
		return err
	}

	if affected != int64(totalTickets) {
		println("database/events/create.go:updateTicket(): affected: " + strconv.FormatInt(affected, 10))
		println("database/events/create.go:updateTicket(): total ticket: " + strconv.Itoa(totalTickets))
		return errors.New("no ticket left to book remains: " + strconv.FormatInt(affected, 10) + ", total book: " + strconv.Itoa(totalTickets))
	}

	return nil
}
