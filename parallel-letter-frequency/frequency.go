package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts runes using parrallel processes.
func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap, len(input))

	for _, w := range input {
		go func(text string) {
			ch <- Frequency(text)
		}(w)
	}

	for range input {
		for k, count := range <-ch {
			m[k] += count
		}
	}

	return m
}
