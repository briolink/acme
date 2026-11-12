package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// BeforeStart is a guard.
//
// Rules:
//   - True when now is strictly before starts_at
func (Impl) BeforeStart(booking acme.Booking, now time.Time) bool {
	panic("not implemented: before_start")
}
