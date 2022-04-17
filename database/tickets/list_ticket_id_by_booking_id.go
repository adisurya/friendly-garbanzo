package tickets

import "github.com/adisurya/friendly-garbanzo/database"

type ListTicketIdByBookingIdResponse struct {
	Id int64 `json:"id"`
}

func ListTicketIdByBookingId(bookingId int64) ([]ListTicketIdByBookingIdResponse, error) {
	db, err := database.Connect()
	tickets := []ListTicketIdByBookingIdResponse{}
	if err != nil {
		println("database/tickets/list_ticket_id_by_booking_id.go:ListTicketIdByBookingId(): " + err.Error())
		return tickets, err
	}
	defer db.Close()

	if err != nil {
		println("database/tickets/list_ticket_id_by_booking_id.go:ListTicketIdByBookingId(): " + err.Error())
		return tickets, err
	}
	rows, err := db.Query("SELECT id FROM tickets WHERE booking_id = ? ORDER BY id ASC", bookingId)
	if err != nil {
		println("database/tickets/list_ticket_id_by_booking_id.go:ListTicketIdByBookingId(): " + err.Error())
		return tickets, err
	}
	defer rows.Close()

	for rows.Next() {
		tick := ListTicketIdByBookingIdResponse{}
		err = rows.Scan(&tick.Id)
		if err != nil {
			println("database/tickets/list_ticket_id_by_booking_id.go:ListTicketIdByBookingId(): " + err.Error())
		} else {
			tickets = append(tickets, tick)
		}
	}
	return tickets, nil
}
