package acme

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// Impl implements acme.Funcs.
type Impl struct{}

var _ acme.Funcs = Impl{}

// rate is the block rate per seat in cents.
func rate(tier acme.Tier) int64 {
	if tier == acme.TierFlexible {
		return 750
	}
	return 600
}

// blocks counts the started 30-minute blocks from a to b.
func blocks(a, b time.Time) int64 {
	d := b.Sub(a)
	n := int64(d / (30 * time.Minute))
	if d%(30*time.Minute) != 0 {
		n++
	}
	return n
}
