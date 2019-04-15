package gigasecond

import "time"

var gs = 1000000000

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(gs))
}
