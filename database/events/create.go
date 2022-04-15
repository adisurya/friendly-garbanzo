package events

import (
	"github.com/adisurya/friendly-garbanzo/database"
	"github.com/adisurya/friendly-garbanzo/structs/requests"
)

func Create(data *requests.CreateEvent) (int64, error) {
	db, err := database.Connect()
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return 0, err
	}
	defer db.Close()

	stmtInst, err := db.Prepare("INSERT INTO events(id, name, event_time, total_tickets, ticket_price, created_at) VALUES(NULL, ?, ?, ?, ?, now())")
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return 0, err
	}
	defer stmtInst.Close()

	result, err := stmtInst.Exec(data.Name, data.EventTime, data.TotalTickets, data.TicketPrice)
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return 0, err
	}

	return result.LastInsertId()
}
