package acme

import (
	"testing"
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// BeforeStart is true exactly when now is strictly before starts_at.
func TestBeforeStart(t *testing.T) {
	starts := time.Date(2026, 3, 10, 10, 0, 0, 0, time.UTC)
	b := acme.Booking{Id: "b1", Guest: "g1", Tier: acme.TierStandard, Seats: 2, StartsAt: starts, EndsAt: starts.Add(2 * time.Hour), Price: 4800}
	tests := []struct {
		name string
		at   time.Duration // now, relative to starts_at
		want bool
	}{
		{"a day before", -24 * time.Hour, true},
		{"one second before", -time.Second, true},
		{"at the start", 0, false},
		{"one second after", time.Second, false},
		{"a day after", 24 * time.Hour, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := (Impl{}).BeforeStart(b, starts.Add(tc.at)); got != tc.want {
				t.Errorf("BeforeStart(now = starts_at %+v) = %v, want %v", tc.at, got, tc.want)
			}
		})
	}
}
