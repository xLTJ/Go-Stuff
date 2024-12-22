// Package gigasecond does gigasecond stuff 😼
package gigasecond

import "time"

// GIGASECOND is the amount of seconds in a "gigasecond"
const GIGASECOND = 1000000000

// AddGigasecond determines the date one gigasecond after a certain date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * GIGASECOND)
}
