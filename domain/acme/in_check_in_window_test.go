package acme

import (
	"testing"
	"time"

	"github.com/briolink/acme/gen/domain/acme"
)

// InCheckInWindow is true from 30 minutes before starts_at to 15 minutes after it, both ends included.
func TestInCheckInWindow(t *testing.T) {
	starts := time.Date(2026, 3, 10, 10, 0, 0, 0, time.UTC)
	b := acme.Booking{Id: "b1", Guest: "g1", Tier: acme.TierStandard, Seats: 2, StartsAt: starts, EndsAt: starts.Add(2 * time.Hour), Price: 4800}
	tests := []struct {
		name string
		at   time.Duration // now, relative to starts_at
		want bool
	}{
		{"an hour before", -time.Hour, false},
		{"one second before the window", -30*time.Minute - time.Second, false},
		{"the window opens", -30 * time.Minute, true},
		{"at the start", 0, true},
		{"the window closes", 15 * time.Minute, true},
		{"one second after the window", 15*time.Minute + time.Second, false},
		{"an hour after", time.Hour, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := (Impl{}).InCheckInWindow(b, starts.Add(tc.at)); got != tc.want {
				t.Errorf("InCheckInWindow(now = starts_at %+v) = %v, want %v", tc.at, got, tc.want)
			}
		})
	}
}
