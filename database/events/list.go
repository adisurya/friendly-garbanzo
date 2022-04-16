package events

import (
	"time"

	"github.com/adisurya/friendly-garbanzo/database"
)

type EventList struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	EventTime    time.Time `json:"event_time"`
	TotalTickets int       `json:"total_tickets"`
	TicketPrice  string    `json:"ticket_price"`
	CreatedAt    time.Time `json:"created_at"`
}

func List() ([]Event, error) {
	db, err := database.Connect()
	events := []Event{}
	if err != nil {
		println("database/events/create.go:List(): " + err.Error())
		return events, err
	}
	defer db.Close()

	if err != nil {
		println("database/events/create.go:List(): " + err.Error())
		return events, err
	}

	rows, err := db.Query("SELECT *FROM events ORDER BY id DESC")
	if err != nil {
		println("database/events/create.go:List(): " + err.Error())
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		evt := Event{}
		err = rows.Scan(&evt.Id, &evt.Name, &evt.EventTime, &evt.TotalTickets, &evt.TicketPrice, &evt.CreatedAt)
		if err != nil {
			println("database/events/create.go:List(): " + err.Error())
		} else {
			events = append(events, evt)
		}
	}
	return events, nil
}
