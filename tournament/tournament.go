package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// contains win and loss of the team
type team struct {
	name                             string
	played, won, drawn, lost, points int
}

// keeps track of the score
type teamsScores map[string]team

// Tally summarizes the competition.
func Tally(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	scores := make(teamsScores)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		if err := scores.addGameScore(line); err != nil {
			return err
		}
	}

	scores.Write(output)

	return nil
}

// Write formats and writes scores, sorted by points.
func (t teamsScores) Write(w io.Writer) {
	fmt.Fprintf(w, "Team%-27s| MP |  W |  D |  L |  P\n", " ")

	teams := allTeams(t)
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})

	for _, team := range teams {
		fmt.Fprintf(w, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			team.name, team.played, team.won, team.drawn, team.lost, team.points)
	}
}

// add game result to overall score
func (t teamsScores) addGameScore(line string) error {
	r := strings.Split(line, ";")
	if len(r) != 3 {
		return fmt.Errorf("invalid line")
	}

	t1, ok := t[r[0]]
	if !ok {
		t1 = team{name: r[0]}
	}

	t2, ok := t[r[1]]
	if !ok {
		t2 = team{name: r[1]}
	}

	err := play(&t1, &t2, r[2])
	if err != nil {
		return err
	}

	t[r[0]], t[r[1]] = t1, t2

	return nil
}

// add game result to the given teams
func play(t1, t2 *team, result string) error {
	t1.played++
	t2.played++

	switch result {
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
		return fmt.Errorf("undefine match result")
	}

	return nil
}

// get slice with all teams and their results
func allTeams(t teamsScores) []team {
	res := make([]team, 0)
	for _, team := range t {
		res = append(res, team)
	}
	return res
}
