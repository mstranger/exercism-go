package robotname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Robot represents a structure for naming robots.
type Robot struct {
	name string
}

// max number of combinations of 2 letters and 3 digits.
const maxNamesCount = 26 * 26 * 10 * 10 * 10

// used names
var names = map[string]bool{}

// Name returns the robot's name if its exists, or generates a new one.
func (r *Robot) Name() (string, error) {
	// robot already has a name
	if r.name != "" {
		return r.name, nil
	}
	// namespace exhausted
	if len(names) == maxNamesCount {
		return "", fmt.Errorf("max namespace")
	}

	var s strings.Builder
	var charCode int

	// generate random name
	t := rand.NewSource(time.Now().UnixNano())
	randSourse := rand.New(t)
	for i := 0; i < 5; i++ {
		if i < 2 {
			// from 'A'..'Z'
			charCode = 65 + randSourse.Intn(26)
		} else {
			// from '0'..'9'
			charCode = 48 + randSourse.Intn(10)
		}
		s.WriteByte(byte(charCode))
	}

	// a robot with the same name already exists,
	// reset the current one and generate it again
	if names[s.String()] {
		r.Reset()
		return r.Name()
	}

	r.name = s.String()
	names[r.name] = true

	return r.name, nil
}

// Reset wipes the current robot's name.
func (r *Robot) Reset() {
	r.name = ""
}
