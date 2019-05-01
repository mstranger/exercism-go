package clock

import "fmt"

// Clock represets a time without dates. Only hours and minutes.
type Clock struct {
	minutes int
	hours   int
}

// minutes per day
const dayMinutes = 24 * 60

// New creates a digital clock.
func New(h, m int) Clock {
	if h < 0 {
		h = 24 + h%24
	}
	if m < 0 {
		m = dayMinutes + m%dayMinutes
	}

	// current time in minutes 10.30 == 630 min
	current := ((h+m/60)*60 + m%60) % dayMinutes
	h, m = current/60, current%60
	return Clock{m, h}
}

// String implements the string representation of the Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.hours, c.minutes)
}

// Add implements adding minutes to the Clock.
func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

// Subtract implements subtracting minutes from the Clock.
func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}
