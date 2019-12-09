package ledger

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Entry represents one record
type Entry struct {
	Date        string
	Description string
	Change      int
}

// for write to channel
type chunk struct {
	i int
	s string
	e error
}

// dict helps print currency according to the lang ("nl-NL", "en-US")
type dict struct {
	sb, delim1, delim2, delim3, se string
}

func defaultEntry() Entry {
	return Entry{Date: "2014-01-01", Description: "", Change: 0}
}

// FormatLedger prints a nicely formatted ledger
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry

	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}
	if len(entries) == 0 {
		_, err := FormatLedger(currency, "en-US", []Entry{defaultEntry()})
		if err != nil {
			return "", err
		}
	}

	sortEntries(entriesCopy)

	s, ok := translateLedger(locale)
	if !ok {
		return "", errors.New("")
	}

	co := make(chan chunk)

	for i, et := range entriesCopy {
		go writeToChan(i, et, locale, currency, co)
	}

	s, err := readFromChan(s, entriesCopy, co)
	if err != nil {
		return "", err
	}

	return s, nil
}

// display records by date in ASC order
func sortEntries(es []Entry) {
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}

	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if checkOrder(m1, m2, first, e) {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}
}

// check records in ASC order
func checkOrder(m1, m2 map[bool]int, e1, e2 Entry) bool {
	return (m1[e2.Date == e1.Date]*m2[e2.Date < e1.Date]*4 +
		m1[e2.Description == e1.Description]*m2[e2.Description < e1.Description]*2 +
		m1[e2.Change == e1.Change]*m2[e2.Change < e1.Change]*1) < 0
}

// output in lang according to the given locale
func translateLedger(locale string) (string, bool) {
	switch locale {
	case "nl-NL":
		return "Datum" +
				strings.Repeat(" ", 10-len("Datum")) +
				" | " + "Omschrijving" +
				strings.Repeat(" ", 25-len("Omschrijving")) +
				" | " + "Verandering" + "\n",
			true
	case "en-US":
		return "Date" +
				strings.Repeat(" ", 10-len("Date")) +
				" | " + "Description" +
				strings.Repeat(" ", 25-len("Description")) +
				" | " + "Change" + "\n",
			true
	default:
		return "", false
	}
}

// read data from the channel
func readFromChan(s string, entries []Entry, co <-chan chunk) (string, error) {
	ss := make([]string, len(entries))
	for range entries {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	for i := 0; i < len(entries); i++ {
		s += ss[i]
	}

	return s, nil
}

// write data to the channel
func writeToChan(i int, entry Entry, locale, currency string, co chan<- chunk) {
	if len(entry.Date) != 10 {
		sendError(co)
	}

	d1, d2, d3, d4, d5 := parseDate(entry.Date)
	if d2 != '-' || d4 != '-' {
		sendError(co)
	}

	de := foldDescription(entry.Description)

	negative := false
	cents := entry.Change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	var a, d string
	switch locale {
	case "nl-NL":
		dt := dict{"", " ", ",", ".", "-"}
		a = joinCurrency(dt, a, cents, negative, currency, co)
		d = d5 + "-" + d3 + "-" + d1
	case "en-US":
		dt := dict{"(", "", ".", ",", ")"}
		a = joinCurrency(dt, a, cents, negative, currency, co)
		d = d3 + "/" + d5 + "/" + d1
	default:
		sendError(co)
	}

	al := utf8.RuneCountInString(a)

	co <- chunk{
		i: i,
		s: d +
			strings.Repeat(" ", 10-len(d)) +
			" | " + de + " | " +
			strings.Repeat(" ", 13-al) + a + "\n",
	}
}

// send error to the channel
func sendError(target chan<- chunk) {
	target <- chunk{e: errors.New("")}
}

// parse given date in format "YYYY-MM-DD"
func parseDate(date string) (string, byte, string, byte, string) {
	return date[0:4], date[4], date[5:7], date[7], date[8:10]
}

// replace long string with <...> sign
func foldDescription(de string) string {
	if len(de) > 25 {
		de = de[:22] + "..."
	} else {
		de = de + strings.Repeat(" ", 25-len(de))
	}

	return de
}

// help function, join formatted currency to the output
func joinCurrency(d dict, a string, cents int, negative bool, currency string, co chan<- chunk) string {
	if currency == "EUR" {
		a += "€"
	} else if currency == "USD" {
		a += "$"
	} else {
		sendError(co)
	}

	a += d.delim1

	centsStr := strconv.Itoa(cents)
	switch len(centsStr) {
	case 1:
		centsStr = "00" + centsStr
	case 2:
		centsStr = "0" + centsStr
	}
	rest := centsStr[:len(centsStr)-2]

	var parts []string

	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}

	for i := len(parts) - 1; i >= 0; i-- {
		a += parts[i] + d.delim3
	}
	a = a[:len(a)-1]

	a += d.delim2
	a += centsStr[len(centsStr)-2:]

	if negative {
		a = d.sb + a + d.se
	} else {
		a += " "
	}

	return a
}
