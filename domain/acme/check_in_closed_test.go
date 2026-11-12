package acme

import (
	"testing"
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// CheckInClosed is true exactly when now is more than 15 minutes after starts_at.
func TestCheckInClosed(t *testing.T) {
	starts := time.Date(2026, 3, 10, 10, 0, 0, 0, time.UTC)
	b := acme.Booking{Id: "b1", Guest: "g1", Tier: acme.TierStandard, Seats: 2, StartsAt: starts, EndsAt: starts.Add(2 * time.Hour), Price: 4800}
	tests := []struct {
		name string
		at   time.Duration // now, relative to starts_at
		want bool
	}{
		{"an hour before", -time.Hour, false},
		{"at the start", 0, false},
		{"the window closes", 15 * time.Minute, false},
		{"one second after the window", 15*time.Minute + time.Second, true},
		{"an hour after", time.Hour, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := (Impl{}).CheckInClosed(b, starts.Add(tc.at)); got != tc.want {
				t.Errorf("CheckInClosed(now = starts_at %+v) = %v, want %v", tc.at, got, tc.want)
			}
		})
	}
}
