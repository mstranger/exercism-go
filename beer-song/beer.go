package beer

import "fmt"

// Verse returns one verse of the song.
func Verse(n int) (string, error) {
	if n > 99 || n < 0 {
		return "", fmt.Errorf("invalid input")
	}

	s, plural := "", "s"

	if n == 0 {
		s = fmt.Sprintf(`No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.
`)
	} else if n == 1 {
		s = fmt.Sprintf(`1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.
`)
	} else {
		if n == 2 {
			plural = ""
		}
		s = fmt.Sprintf(`%d bottles of beer on the wall, %[1]d bottles of beer.
Take one down and pass it around, %d bottle%s of beer on the wall.
`, n, n-1, plural)
	}

	return s, nil
}

// Verses returns a piece from start to end.
func Verses(start, end int) (string, error) {
	if start < end {
		return "", fmt.Errorf("start < end")
	}

	s := ""
	for i := start; i >= end; i-- {
		line, err := Verse(i)
		if err != nil {
			return "", err
		}

		s += line + "\n"
	}

	return s, nil
}

// Song returns all song "99 Bottles of Beer on the Wall".
func Song() string {
	s, _ := Verses(99, 0)
	return s
}
