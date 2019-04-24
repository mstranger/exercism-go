package scale

import "strings"

// Scale generates the musical scale starting with the tonic
// and following the specified interval pattern.
func Scale(tonic string, interval string) []string {
	var scale []string

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}

	tonic = strings.Title(tonic)
	for i, elem := range scale {
		if elem == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}

	partial := []string{}
	steps := map[string]int{"m": 1, "M": 2, "A": 3}

	var i int
	for _, v := range strings.Split(interval, "") {
		if step, ok := steps[v]; ok {
			partial = append(partial, scale[i%len(scale)])
			i += step
		}
	}

	return partial
}
