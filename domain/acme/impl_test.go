package acme

import (
	"testing"

	"github.com/briolink/acme/gen/domain/acme/acmetest"
	"github.com/briolink/acme/samples"
)

func TestBooking(t *testing.T) { acmetest.Booking(t, Impl{}, samples.Booking{}) }
