package kindergarten

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

// Garden lists plants owned by students.
type Garden map[string][]string

var plants = map[byte]string{
	'C': "clover",
	'R': "radishes",
	'G': "grass",
	'V': "violets",
}

// NewGarden creates a new kindergarten garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	if diagram[0] != '\n' {
		return nil, fmt.Errorf("invalid diagram")
	}

	g := &Garden{}
	scanner := bufio.NewScanner(strings.NewReader(diagram[1:]))

	childrenDup := make([]string, len(children))
	copy(childrenDup, children)
	sort.Strings(childrenDup)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 2*len(children) {
			return nil, fmt.Errorf("invalid rows")
		}

		j := 0
		for i := 0; i < len(line); i++ {
			if _, ok := plants[line[i]]; !ok {
				return nil, fmt.Errorf("invalid cups")
			}
			if j > 0 && childrenDup[j] == childrenDup[j-1] {
				return nil, fmt.Errorf("duplicate name")
			}

			(*g)[childrenDup[j]] = append((*g)[childrenDup[j]], plants[line[i]])

			if i%2 != 0 {
				j++
			}
		}
	}

	return g, nil
}

// Plants lists the plants owned by a child in the garden.
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
