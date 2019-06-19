package house

import (
	"fmt"
	"strings"
)

var actors = []struct {
	name   string
	action string
}{
	{"house", "Jack built"},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

// Verse plays one verse for the given number.
func Verse(n int) string {
	s := "This is "
	for i := range actors[:n] {
		s += fmt.Sprintf("the %s\nthat %s ",
			actors[n-i-1].name, actors[n-i-1].action)
	}

	idx := strings.LastIndex(s, "\n")
	s = s[:idx] + " " + s[idx+1:] // del the last newline
	s = s[:len(s)-1] + "."        // del the last space

	return s
}

// Song returns the liric "This is the House that Jack Built".
func Song() string {
	var b strings.Builder
	for i := 1; i <= 12; i++ {
		b.WriteString(Verse(i))
		b.WriteString("\n\n")
	}

	return strings.TrimSpace(b.String())
}
