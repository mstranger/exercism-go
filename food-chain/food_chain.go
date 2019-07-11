package foodchain

import "strings"

const TestVerstion = 1

var actors = []struct {
	name  string
	aside string
}{
	{"fly", ""},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

// Verse returns the text of the given verse.
func Verse(n int) string {
	if n < 0 || n > 8 {
		return ""
	}

	s := "I know an old lady who swallowed a " + actors[n-1].name + ".\n"
	s += actors[n-1].aside
	if actors[n-1].aside != "" {
		s += "\n"
	}
	s += middle(n)

	return strings.TrimSpace(s)
}

// Verses outputs lyrics from a to b verse.
func Verses(a, b int) string {
	var s strings.Builder
	for i := a; i < b; i++ {
		s.WriteString(Verse(i))
		s.WriteString("\n\n")
	}
	s.WriteString(Verse(b))
	return s.String()
}

// Song prints all song.
func Song() string {
	return Verses(1, 8)
}

func middle(n int) string {
	if n == 0 {
		return "I don't know why she swallowed the fly. Perhaps she'll die."
	}

	var s string
	if n > 1 {
		// the last verse
		if n == 8 {
			return s
		}

		s = "She swallowed the " + actors[n-1].name + " to catch the " +
			actors[n-2].name + ".\n"

		// string with the spider
		if n == 3 {
			t := actors[1].aside
			s = s[:len(s)-2] + " that" + t[2:] + "\n"
		}
	}

	s += middle(n - 1)

	return s
}
