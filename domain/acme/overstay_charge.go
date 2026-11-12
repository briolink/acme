package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// OverstayCharge is a pure.
// What a check-out at now costs on top of the price
//
// Rules:
//   - 0 when now is no later than 10 minutes after ends_at
//   - Otherwise the seats times the blocks from ends_at to now times the tier's block rate, as in quote; the grace minutes count toward the first block
func (Impl) OverstayCharge(booking acme.Booking, now time.Time) int64 {
	if !now.After(booking.EndsAt.Add(10 * time.Minute)) {
		return 0
	}
	return booking.Seats * blocks(booking.EndsAt, now) * rate(booking.Tier)
}
