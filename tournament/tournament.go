// package tournament

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team respresents team name.
type Team string

// Matches respresents matches played.
type Matches struct {
	played, won, drawn, lost, points int
}

// TeamResult is a result for one command.
type TeamResult struct {
	Team
	Matches
}

// Results keeps track of the score.
type Results map[Team]Matches

// Tally summarizes the competition.
func Tally(input io.Reader, output io.Writer) error {
	// reader := bufio.NewReader(input)
	table := NewTable()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		if err := addToResultsTable(line, table); err != nil {
			return err
		}
	}

	output.Write([]byte(table.String()))

	/*
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF && line == "" {
				break
			}

			line = strings.TrimSpace(line)
			// skip comments or empty lines
			if strings.HasPrefix(line, "#") || line == "" {
				continue
			}

			if err := addToResultsTable(line, table); err != nil {
				return err
			}
		}

		output.Write([]byte(table.String()))
	*/

	return nil
}

// NewTable creates a new results table.
func NewTable() *Results {
	return &Results{}
}

// String implements Stringer interface for Results type.
func (r *Results) String() string {
	s := fmt.Sprintf("Team%-27s| MP |  W |  D |  L |  P\n", " ")

	t := toArray(r)
	// sort array by points
	sort.Slice(t, func(i, j int) bool {
		if t[i].points == t[j].points {
			return t[i].Team < t[j].Team
		}
		return t[i].points > t[j].points
	})

	for _, v := range t {
		s1 := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",
			v.Team, v.played, v.won, v.drawn, v.lost, v.points)
		s += s1
	}

	return s
}

// Join combines two results.
func (r *Results) Join(t *Results) {
	for k, v := range *t {
		matches := (*r)[k]
		matches.played += v.played
		matches.won += v.won
		matches.lost += v.lost
		matches.drawn += v.drawn
		matches.points += v.points
		(*r)[k] = matches
	}
}

func addToResultsTable(line string, t *Results) error {
	r := strings.Split(line, ";")
	if len(r) != 3 {
		return fmt.Errorf("invalid line")
	}

	t1 := Matches{played: 1}
	t2 := Matches{played: 1}

	switch r[2] {
	case "win":
		t1.won++
		t1.points += 3
		t2.lost++
	case "loss":
		t2.won++
		t2.points += 3
		t1.lost++
	case "draw":
		t1.drawn++
		t1.points++
		t2.drawn++
		t2.points++
	default:
		return fmt.Errorf("undefined result")
	}

	// curent match
	m := Results{
		Team(r[0]): t1,
		Team(r[1]): t2,
	}

	// join current result with results table
	t.Join(&m)

	return nil
}

func toArray(r *Results) []TeamResult {
	res := make([]TeamResult, 0)

	for k, v := range *r {
		t := TeamResult{
			Team:    k,
			Matches: v,
		}
		res = append(res, t)
	}

	return res
}

func main() {
	input := `
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
`

	var out bytes.Buffer

	Tally(strings.NewReader(input), &out)

	fmt.Println(out.String())
}
