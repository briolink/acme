package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// CheckOut is a action.
//
// Rules:
//   - Sets checked_out_at to now and overstay_fee to overstay_charge of the booking at now; nothing else changes
func (Impl) CheckOut(booking acme.Booking, now time.Time) acme.Booking {
	booking.CheckedOutAt = &now
	booking.OverstayFee = Impl{}.OverstayCharge(booking, now)
	return booking
}
