package acme

import (
	"github.com/briolink/acme/gen/domain/acme"
)

// MoneyWithinBounds is a invariant.
//
// Rules:
//   - True when the deposit is between 0 and the price, the refund is between 0 and the deposit, and the overstay fee is not negative; every bound is included
func (Impl) MoneyWithinBounds(booking acme.Booking) bool {
	return booking.Deposit >= 0 && booking.Deposit <= booking.Price &&
		booking.Refund >= 0 && booking.Refund <= booking.Deposit && booking.OverstayFee >= 0
}
