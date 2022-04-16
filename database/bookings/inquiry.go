package bookings

import (
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
)

type BookingInquiry struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	EventId     int       `json:"event_id"`
	TotalBooked int       `json:"total_booked"`
	TicketPrice string    `json:"ticket_price"`
	TotalPrice  string    `json:"total_price"`
	ValidUntil  time.Time `json:"valid_until"`
	IsPaid      int       `json:"is_paid"`
	CreatedAt   time.Time `json:"created_at"`
	EventName   string    `json:"event_name"`
	EventTime   time.Time `json:"event_time"`
}

func Inquiry(id int64) (BookingInquiry, error) {
	db, err := database.Connect()
	inquiry := BookingInquiry{}

	if err != nil {
		println("database/bookings/inquiry.go:Inquiry(): " + err.Error())
		return inquiry, err
	}
	defer db.Close()

	if err != nil {
		println("database/bookings/inquiry.go:Inquiry(): " + err.Error())
		return inquiry, err
	}

	row := db.QueryRow(`SELECT b.*,
		e.name AS event_name , e.event_time
		FROM bookings b
		INNER JOIN events e ON(b.event_id = e.id)
		WHERE b.id = ? AND b.is_paid = 0`, id)
	err = row.Scan(
		&inquiry.Id, &inquiry.Name, &inquiry.EventId, &inquiry.TotalBooked,
		&inquiry.TicketPrice, &inquiry.TotalPrice, &inquiry.ValidUntil,
		&inquiry.IsPaid, &inquiry.CreatedAt, &inquiry.EventName, &inquiry.EventTime)
	if err != nil {
		println("database/bookings/inquiry.go:Inquiry(): " + err.Error())
		return inquiry, err
	}

	return inquiry, nil
}
