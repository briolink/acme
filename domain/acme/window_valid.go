package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// WindowValid is a invariant.
//
// Rules:
//   - True when ends_at is strictly after starts_at and the booking lasts at most 8 hours
func (Impl) WindowValid(booking acme.Booking) bool {
	return booking.EndsAt.After(booking.StartsAt) && booking.EndsAt.Sub(booking.StartsAt) <= 8*time.Hour
}
