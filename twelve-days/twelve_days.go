package twelve

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

const tmpl = `
On the {{.data.DayString}} day of Christmas my true love gave to me: {{listAll .idx}} in a Pear Tree.
`

type twelveDays struct {
	DayString, GiftCount, Gift string
}

var gifts = []twelveDays{
	{"first", "a", "Partridge"},
	{"second", "two", "Turtle Doves"},
	{"third", "three", "French Hens"},
	{"fourth", "four", "Calling Birds"},
	{"fifth", "five", "Gold Rings"},
	{"sixth", "six", "Geese-a-Laying"},
	{"seventh", "seven", "Swans-a-Swimming"},
	{"eighth", "eight", "Maids-a-Milking"},
	{"ninth", "nine", "Ladies Dancing"},
	{"tenth", "ten", "Lords-a-Leaping"},
	{"eleventh", "eleven", "Pipers Piping"},
	{"twelfth", "twelve", "Drummers Drumming"},
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

	tmpl, err := template.New("twelve").Funcs(template.FuncMap{
		"listAll": giftsOnDay,
	}).Parse(tmpl)
	if err != nil {
		log.Println(err)
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, map[string]interface{}{
		"idx":  n,
		"data": gifts[n-1],
	})
	if err != nil {
		log.Println(err)
	}

	return strings.TrimSpace(b.String())
}

// return a list of all gifts for the day
func giftsOnDay(n int) string {
	if n == 1 {
		return gifts[0].GiftCount + " " + gifts[0].Gift
	}

	return giftsExceptFirst(n) + "and " + giftsOnDay(1)
}

// return all gifts for the day except first
func giftsExceptFirst(n int) string {
	if n == 1 {
		return ""
	}

	return strings.Join([]string{gifts[n-1].GiftCount + " " + gifts[n-1].Gift,
		giftsExceptFirst(n - 1)}, ", ")
}
