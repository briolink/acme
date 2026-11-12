package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// RefundDue is a pure.
// What a cancellation at now returns
//
// Rules:
//   - 0 when the deposit is 0
//   - Flexible tier: the whole deposit when now is at least 24 hours before starts_at, otherwise half the deposit rounded down
//   - Standard tier: the whole deposit when now is at least 72 hours before starts_at, half the deposit rounded down when at least 24 hours before, otherwise 0
func (Impl) RefundDue(booking acme.Booking, now time.Time) int64 {
	if booking.Deposit == 0 {
		return 0
	}
	until := booking.StartsAt.Sub(now)
	if booking.Tier == acme.TierFlexible {
		if until >= 24*time.Hour {
			return booking.Deposit
		}
		return booking.Deposit / 2
	}
	switch {
	case until >= 72*time.Hour:
		return booking.Deposit
	case until >= 24*time.Hour:
		return booking.Deposit / 2
	}
	return 0
}
