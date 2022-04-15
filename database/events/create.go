package events

import (
	"database/sql"

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

	trx, err := db.Begin()
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return 0, err
	}

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

	insertId, err := result.LastInsertId()
	if err != nil {
		println("database/events/create.go:Create(): " + err.Error())
		return 0, err
	}

	errT := createTicket(insertId, data.TotalTickets, db)
	if errT != nil {
		println("database/events/create.go:Create(): " + errT.Error())

		trx.Rollback()

		return 0, err
	}

	errC := trx.Commit()
	if errC != nil {
		println("database/events/create.go:Create(): " + errC.Error())

		return 0, err
	}

	return insertId, nil
}

func createTicket(id int64, totalTickets int, db *sql.DB) error {

	stmtInst, err := db.Prepare("INSERT INTO tickets(id, event_id, created_at) VALUES(NULL, ?, now())")
	if err != nil {
		println("database/events/create.go:createTicket(): " + err.Error())
		return err
	}
	defer stmtInst.Close()

	for i := 0; i < totalTickets; i++ {
		_, err := stmtInst.Exec(id)
		if err != nil {
			println("database/events/create.go:createTicket(): " + err.Error())
			return err
		}

	}

	return nil
}
