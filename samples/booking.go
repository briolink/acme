// Package samples is the hand-written half of the acme derived tests: the
// entities and events at the boundaries the rules in spec/domain/acme.json
// name. It is not generated.
package samples

import (
	"time"

	"github.com/briolink/acme/gen/domain/acme"
	"github.com/briolink/acme/gen/domain/acme/acmetest"
)

// starts is the slot every sample books: two hours, four blocks, two seats on
// the standard tier, so the price is 4800 cents and the minimum deposit 960.
var starts = time.Date(2026, 3, 10, 10, 0, 0, 0, time.UTC)

func booking(deposit int64) acme.Booking {
	return acme.Booking{
		Id: "b1", Guest: "g1", Tier: acme.TierStandard, Seats: 2,
		StartsAt: starts, EndsAt: starts.Add(2 * time.Hour),
		Price: 4800, Deposit: deposit,
	}
}

type c = acmetest.BookingCase

func bc(e acme.Booking, ev acme.BookingEvent) c { return c{Entity: e, Event: ev} }

// Booking implements acmetest.BookingSamples.
type Booking struct{}

func confirm(deposit int64, now time.Time) acme.BookingConfirm {
	return acme.BookingConfirm{Deposit: deposit, Now: now}
}

// Fire: transitions in the order of the spec; each rule's boundary from the
// inside.
func (Booking) Fire(i int) []c {
	switch i {
	case 0: // requested confirm: the minimum deposit, the price, one second before the start
		return []c{
			bc(booking(0), confirm(960, starts.Add(-time.Hour))),
			bc(booking(0), confirm(4800, starts.Add(-time.Second))),
		}
	case 1: // requested cancel: no guard, nothing was paid
		return []c{bc(booking(0), acme.BookingCancel{Now: starts.Add(-time.Hour)})}
	case 2: // confirmed cancel: before the start, with the refund at each threshold
		return []c{
			bc(booking(4800), acme.BookingCancel{Now: starts.Add(-time.Second)}),
			bc(booking(4800), acme.BookingCancel{Now: starts.Add(-24 * time.Hour)}),
			bc(booking(4800), acme.BookingCancel{Now: starts.Add(-72 * time.Hour)}),
		}
	case 3: // confirmed check_in: both ends of the window are included
		return []c{
			bc(booking(4800), acme.BookingCheckIn{Now: starts.Add(-30 * time.Minute)}),
			bc(booking(4800), acme.BookingCheckIn{Now: starts}),
			bc(booking(4800), acme.BookingCheckIn{Now: starts.Add(15 * time.Minute)}),
		}
	case 4: // confirmed mark_no_show: one second after the window closes, and long after
		return []c{
			bc(booking(4800), acme.BookingMarkNoShow{Now: starts.Add(15*time.Minute + time.Second)}),
			bc(booking(4800), acme.BookingMarkNoShow{Now: starts.Add(5 * time.Hour)}),
		}
	case 5: // checked_in check_out: on time, at the end of the grace, one second over, two blocks over
		in := booking(4800)
		at := starts
		in.CheckedInAt = &at
		end := in.EndsAt
		return []c{
			bc(in, acme.BookingCheckOut{Now: end}),
			bc(in, acme.BookingCheckOut{Now: end.Add(10 * time.Minute)}),
			bc(in, acme.BookingCheckOut{Now: end.Add(10*time.Minute + time.Second)}),
			bc(in, acme.BookingCheckOut{Now: end.Add(40 * time.Minute)}),
		}
	}
	panic("no such transition")
}

// Block: each guard false, one step outside each boundary.
func (Booking) Block(i int) []c {
	switch i {
	case 0: // one cent below the minimum, one above the price, at the start, after it
		return []c{
			bc(booking(0), confirm(959, starts.Add(-time.Hour))),
			bc(booking(0), confirm(4801, starts.Add(-time.Hour))),
			bc(booking(0), confirm(960, starts)),
			bc(booking(0), confirm(960, starts.Add(time.Second))),
		}
	case 2: // at the start, one second after
		return []c{
			bc(booking(4800), acme.BookingCancel{Now: starts}),
			bc(booking(4800), acme.BookingCancel{Now: starts.Add(time.Second)}),
		}
	case 3: // one second before the window opens, one second after it closes
		return []c{
			bc(booking(4800), acme.BookingCheckIn{Now: starts.Add(-30*time.Minute - time.Second)}),
			bc(booking(4800), acme.BookingCheckIn{Now: starts.Add(15*time.Minute + time.Second)}),
		}
	case 4: // exactly when the window closes, and well before it
		return []c{
			bc(booking(4800), acme.BookingMarkNoShow{Now: starts.Add(15 * time.Minute)}),
			bc(booking(4800), acme.BookingMarkNoShow{Now: starts.Add(-time.Hour)}),
		}
	}
	panic("no guard on this transition")
}

// Forbid: an entity of the state and the event it must reject.
func (Booking) Forbid(state acme.BookingState, event string) c {
	e := booking(0)
	if state != acme.BookingStateRequested {
		e.Deposit = 4800
	}
	now := starts.Add(-time.Hour)
	switch event {
	case "confirm":
		return bc(e, confirm(960, now))
	case "cancel":
		return bc(e, acme.BookingCancel{Now: now})
	case "check_in":
		return bc(e, acme.BookingCheckIn{Now: now})
	case "mark_no_show":
		return bc(e, acme.BookingMarkNoShow{Now: now})
	case "check_out":
		return bc(e, acme.BookingCheckOut{Now: now})
	}
	panic("no such event")
}
