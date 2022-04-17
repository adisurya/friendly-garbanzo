package requests

type CreatePayment struct {
	BookingId int64  `json:"booking_id"  validate:"required"`
	Total     string `json:"total" validate:"required"`
}
