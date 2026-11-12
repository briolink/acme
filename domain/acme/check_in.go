package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// CheckIn is a action.
//
// Rules:
//   - Sets checked_in_at to now; nothing else changes
func (Impl) CheckIn(booking acme.Booking, now time.Time) acme.Booking {
	booking.CheckedInAt = &now
	return booking
}
