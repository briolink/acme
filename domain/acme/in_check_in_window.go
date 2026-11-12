package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// InCheckInWindow is a guard.
//
// Rules:
//   - True when now is no earlier than 30 minutes before starts_at and no later than 15 minutes after it; both ends are included
func (Impl) InCheckInWindow(booking acme.Booking, now time.Time) bool {
	return !now.Before(booking.StartsAt.Add(-30*time.Minute)) && !now.After(booking.StartsAt.Add(15*time.Minute))
}
