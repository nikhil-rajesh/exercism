// Package clock implements a simple clock.
package clock

import "fmt"

// Clock stores hours and minutes of a clock.
type Clock struct {
	min int
}

var dayMins = 24 * 60

// Add adds minutes to the clock
func (c Clock) Add(x int) Clock {
	return New(0, c.min+x)
}

// Subtract removes minutes from the clock
func (c Clock) Subtract(x int) Clock {
	return New(0, c.min-x)
}

// String returns the time in clock in string form
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

// New returns a new Clock
func New(h int, m int) Clock {
	total := (h*60 + m) % dayMins
	if total < 0 {
		total += dayMins
	}
	return Clock{total}
}
