package requests

type CreateEvent struct {
	Name         string `json:"name" validate:"required"`
	EventTime    string `json:"event_time"  validate:"required"`
	TotalTickets int    `json:"total_tickets" validate:"required"`
	TicketPrice  int    `json:"ticket_price" validate:"required"`
}
