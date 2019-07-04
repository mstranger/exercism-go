package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

type twelveDays struct {
	DayString, Gift string
}

var gifts = []twelveDays{
	{"first", "a Partridge"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song outputs the lirics "The Twelve Days of Christmas".
func Song() string {
	var b bytes.Buffer
	for i := 1; i <= 12; i++ {
		b.WriteString(Verse(i))
		b.WriteString("\n")
	}
	return b.String()
}

// Verse returns one line for the given day.
func Verse(n int) string {
	if n < 1 || n > 12 {
		return ""
	}

	s := fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me: %s in a Pear Tree.",
		gifts[n-1].DayString, giftsOnDay(n))

	if n == 1 {
		s = strings.Replace(s, " and", "", -1)
	}

	return s
}

// return a list of all gifts for the day
func giftsOnDay(n int) string {
	if n == 1 {
		return "and " + gifts[0].Gift
	}

	return strings.Join([]string{gifts[n-1].Gift, giftsOnDay(n - 1)}, ", ")
}
