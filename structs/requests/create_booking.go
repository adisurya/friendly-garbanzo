package requests

type CreateBooking struct {
	Name    string `json:"name" validate:"required"`
	EventId int64  `json:"event_id"  validate:"required"`
	Total   int    `json:"total" validate:"required"`
}
