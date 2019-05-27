package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a structure for naming robots.
type Robot struct {
	name string
}

// max number of combinations of 2 letters and 3 digits
const maxNamesCount = 26 * 26 * 10 * 10 * 10

// used names
var names = map[string]bool{}

// not necessary
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name returns the robot's name if its exists, or generates a new one.
func (r *Robot) Name() (string, error) {
	// robot already has a name
	if r.name != "" {
		return r.name, nil
	}
	// namespace exhausted
	if len(names) >= maxNamesCount {
		return "", fmt.Errorf("max namespace")
	}

	for r.name = name(); names[r.name]; {
		r.name = name()
	}

	names[r.name] = true

	return r.name, nil
}

// Reset wipes the current robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

// generate a random name, such as `AB123`.
func name() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	n := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, n)
}
