package acme

import (
	"github.com/briolink/acme/gen/domain/acme"
)

// Confirm is a action.
//
// Rules:
//   - Sets deposit; nothing else changes
func (Impl) Confirm(booking acme.Booking, deposit int64) acme.Booking {
	booking.Deposit = deposit
	return booking
}
