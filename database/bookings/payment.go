package bookings

import (
	"errors"

	"github.com/adisurya/friendly-garbanzo/database"
	"github.com/adisurya/friendly-garbanzo/database/tickets"
)

type PaymentResponse struct {
	Booking BookingInquiry                            `json:"booking"`
	Tickets []tickets.ListTicketIdByBookingIdResponse `json:"tickets"`
}

func Payment(bookingId int64) (PaymentResponse, error) {
	res := PaymentResponse{}

	db, err := database.Connect()
	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}
	defer db.Close()

	stmtInst, err := db.Prepare(`UPDATE bookings SET is_paid = 1 WHERE id = ?`)
	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}
	defer stmtInst.Close()

	result, err := stmtInst.Exec(bookingId)
	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}

	if affected != 1 {
		return res, errors.New("Failed update booking data")
	}
	res.Booking, err = Detail(bookingId)

	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}

	res.Tickets, err = tickets.ListTicketIdByBookingId(bookingId)

	if err != nil {
		println("database/bookings/payment.go:Payment(): " + err.Error())
		return res, err
	}
	return res, nil

}
