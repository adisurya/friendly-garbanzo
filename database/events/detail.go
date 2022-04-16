package events

import (
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
)

type EventDetail struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	EventTime     time.Time `json:"event_time"`
	TotalTickets  int       `json:"total_tickets"`
	TicketBooked  int       `json:"ticket_booked"`
	TicketRemains int       `json:"ticket_remains"`
	TicketPrice   string    `json:"ticket_price"`
	CreatedAt     time.Time `json:"created_at"`
}

func Detail(id int64) (EventDetail, error) {
	db, err := database.Connect()
	detail := EventDetail{}

	if err != nil {
		println("database/events/create.go:Detail(): " + err.Error())
		return detail, err
	}
	defer db.Close()

	if err != nil {
		println("database/events/create.go:Detail(): " + err.Error())
		return detail, err
	}

	row := db.QueryRow(`
	SELECT e.*,
		SUM(case when t.booking_id IS NULL	THEN 0 ELSE 1 END) AS booked,
		SUM(case when t.booking_id IS NULL	THEN 1 ELSE 0 END) AS remains

	FROM events e
	LEFT JOIN tickets t ON (e.id = t.event_id)
	WHERE e.id = ?`, id)
	err = row.Scan(&detail.Id, &detail.Name, &detail.EventTime, &detail.TotalTickets, &detail.TicketPrice, &detail.CreatedAt, &detail.TicketBooked, &detail.TicketRemains)
	if err != nil {
		println("database/events/create.go:Detail(): " + err.Error())
		return detail, err
	}

	return detail, nil
}
