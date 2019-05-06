package meetup

import "time"

// WeekSchedule represents int number.
type WeekSchedule int

// Values for WeekSchedule.
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day calculates the day of meetups.
func Day(week WeekSchedule, day time.Weekday, month time.Month, year int) int {
	var dayNumber int
	switch {
	case week == Last:
		// if Feb
		n := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
		if n == 28 {
			dayNumber = 1 + 6*3
		} else {
			dayNumber = 1 + 6*4
		}
	case week == Teenth:
		dayNumber = 13
	default:
		dayNumber = int(1 + 7*week)
	}

	t := time.Date(year, month, dayNumber, 0, 0, 0, 0, time.UTC)

	for t.Weekday() != day {
		t = t.Add(time.Hour * 24)
	}

	return t.Day()
}
