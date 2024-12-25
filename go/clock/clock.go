package clock

import "fmt"

const (
	minutePerHour = 60
	minutePerDay  = 1440
)

// Clock is a type for defining a clock
type Clock struct {
	hours   int
	minutes int
}

// New creates a new Clock
func New(h, m int) Clock {
	totalMinutes := (m + h*minutePerHour) % minutePerDay
	if totalMinutes < 0 {
		totalMinutes += minutePerDay
	}

	return Clock{hours: totalMinutes / minutePerHour, minutes: totalMinutes % minutePerHour}
}

// Add adds a number of minutes to a Clock
func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

// Subtract subtracts a number of minutes from a Clock
//
// Within the if-statements, first modulo is used to first get the remainder, then make it positive by adding the base.
// Second modulo is used to account for cases with 24 hours or 60 minutes
func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
