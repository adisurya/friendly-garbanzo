package tickets

import (
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
)

type TicketDetailStruct struct {
	Id          int64      `json:"id"`
	EventId     int64      `json:"event_id"`
	BookingId   *int64     `json:"booking_id"`
	CreatedAt   time.Time  `json:"created_at"`
	Name        *string    `json:"name"`
	TicketPrice string     `json:"ticket_price"`
	IsPaid      *int       `json:"is_paid"`
	ValidUntil  *time.Time `json:"valid_until"`
	EventName   string     `json:"event_name"`
	EventTime   time.Time  `json:"event_time"`
}

func Detail(id int64) (TicketDetailStruct, error) {
	db, err := database.Connect()
	detail := TicketDetailStruct{}

	if err != nil {
		println("database/tickets/detail.go:Detail(): " + err.Error())
		return detail, err
	}
	defer db.Close()

	row := db.QueryRow(`
	SELECT 
		t.*, b.name,
		(case when b.ticket_price IS NULL THEN e.ticket_price ELSE b.ticket_price END) as ticket_price
		, b.is_paid, b.valid_until,
		e.name AS event_name , e.event_time
	FROM tickets t 
	LEFT JOIN bookings b ON(t.booking_id = b.id)
	INNER JOIN events e ON(t.event_id = e.id)
	WHERE t.id = ?
	`, id)
	err = row.Scan(&detail.Id, &detail.EventId, &detail.BookingId, &detail.CreatedAt, &detail.Name, &detail.TicketPrice, &detail.IsPaid, &detail.ValidUntil, &detail.EventName, &detail.EventTime)
	if err != nil {
		println("database/tickets/detail.go:Detail(): " + err.Error())
		return detail, err
	}

	return detail, nil

}
