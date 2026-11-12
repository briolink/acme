package acme

import (
	"testing"
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// DepositCovers is true when the deposit is at least the minimum (20% of the
// price, rounded up to the next cent) and at most the price, and now is
// strictly before starts_at.
func TestDepositCovers(t *testing.T) {
	starts := time.Date(2026, 3, 10, 10, 0, 0, 0, time.UTC)
	booking := func(price int64) acme.Booking {
		return acme.Booking{Id: "b1", Guest: "g1", Tier: acme.TierStandard, Seats: 1, StartsAt: starts, EndsAt: starts.Add(time.Hour), Price: price}
	}
	tests := []struct {
		name    string
		price   int64
		deposit int64
		at      time.Duration // now, relative to starts_at
		want    bool
	}{
		{"the minimum", 4800, 960, -time.Hour, true},
		{"one cent below the minimum", 4800, 959, -time.Hour, false},
		{"the price", 4800, 4800, -time.Hour, true},
		{"one cent above the price", 4800, 4801, -time.Hour, false},
		{"between the bounds", 4800, 2000, -time.Hour, true},
		{"zero", 4800, 0, -time.Hour, false},
		{"a negative deposit", 4800, -1, -time.Hour, false},
		{"a minimum that rounds up", 4801, 961, -time.Hour, true},
		{"one cent below a minimum that rounds up", 4801, 960, -time.Hour, false},
		{"a free booking with no deposit", 0, 0, -time.Hour, true},
		{"one second before the start", 4800, 960, -time.Second, true},
		{"at the start", 4800, 960, 0, false},
		{"one second after the start", 4800, 960, time.Second, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := (Impl{}).DepositCovers(booking(tc.price), tc.deposit, starts.Add(tc.at)); got != tc.want {
				t.Errorf("DepositCovers(price %d, deposit %d, now = starts_at %+v) = %v, want %v", tc.price, tc.deposit, tc.at, got, tc.want)
			}
		})
	}
}
