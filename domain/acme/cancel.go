package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// Cancel is a action.
//
// Rules:
//   - Sets cancelled_at to now and refund to refund_due of the booking at now; nothing else changes
func (Impl) Cancel(booking acme.Booking, now time.Time) acme.Booking {
	booking.CancelledAt = &now
	booking.Refund = Impl{}.RefundDue(booking, now)
	return booking
}
