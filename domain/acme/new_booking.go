package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// NewBooking is a pure.
// Creates a requested booking
//
// Rules:
//   - Returns a booking with the given values and its price from quote; deposit, refund and overstay fee are 0 and no timestamp is set
func (Impl) NewBooking(id string, guest string, tier acme.Tier, seats int64, starts_at time.Time, ends_at time.Time) acme.Booking {
	return acme.Booking{Id: id, Guest: guest, Tier: tier, Seats: seats, StartsAt: starts_at, EndsAt: ends_at,
		Price: Impl{}.Quote(tier, seats, starts_at, ends_at)}
}
