package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// DepositCovers is a guard.
//
// Rules:
//   - True when the deposit is at least the minimum deposit and at most the price, and now is strictly before starts_at
//   - The minimum deposit is 20% of the price, rounded up to the next cent
func (Impl) DepositCovers(booking acme.Booking, deposit int64, now time.Time) bool {
	minimum := (booking.Price*20 + 99) / 100
	return deposit >= minimum && deposit <= booking.Price && now.Before(booking.StartsAt)
}
