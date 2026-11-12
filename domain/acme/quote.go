package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// Quote is a pure.
// The price of a slot
//
// Rules:
//   - The block rate per seat is 600 cents on the standard tier and 750 on the flexible tier
//   - The price is the seats times the blocks from starts_at to ends_at times the block rate; a block is a started 30 minutes
func (Impl) Quote(tier acme.Tier, seats int64, starts_at time.Time, ends_at time.Time) int64 {
	return seats * blocks(starts_at, ends_at) * rate(tier)
}
