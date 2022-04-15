package requests

type CreateEvent struct {
	Name         string `json:"name"`
	EventTime    string `json:"event_time"`
	TotalTickets int    `json:"total_ticket"`
	TicketPrice  int    `json:"ticket_price"`
}
