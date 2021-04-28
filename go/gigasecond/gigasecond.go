package gigasecond

import "time"

// Adds a Gigasecond to the input time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
