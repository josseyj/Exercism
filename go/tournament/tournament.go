package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const winPoint int = 3
const lossPoint int = 0
const drawPoint int = 1

var actions = map[string]func(*Score, *Score){
	"win":  func(team1 *Score, team2 *Score) { team1.wins++; team2.looses++ },
	"loss": func(team1 *Score, team2 *Score) { team1.looses++; team2.wins++ },
	"draw": func(team1 *Score, team2 *Score) { team1.draws++; team2.draws++ },
}

// ------------------------------------------------------------------------

//Score type deines a team score board.
type Score struct {
	team   string
	wins   int
	looses int
	draws  int
}

//Matches returns the number of matches played
func (s Score) matches() int {
	return s.wins + s.looses + s.draws
}

//Points returns the points for the team
func (s Score) points() int {
	return s.wins*winPoint + s.looses*lossPoint + s.draws*drawPoint
}

// ------------------------------------------------------------------------

//Scores type is an array of team scores
type Scores []*Score

// Len returns the length of Scores type
func (s Scores) Len() int {
	return len(s)
}

// Less indicates if the item at 'i' is less that the item at index 'j'.
func (s Scores) Less(i, j int) bool {
	return s[i].points() > s[j].points() ||
		(s[i].points() == s[j].points() && s[i].team < s[j].team)
}

// Swap swaps the elements with indexes i and j.
func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ------------------------------------------------------------------------

// Tally reads the input and build a score board
func Tally(reader io.Reader, writer io.Writer) error {
	teams := map[string]*Score{}
	inputBuffer := bufio.NewReader(reader)
	for {
		line, err := inputBuffer.ReadString('\n')
		line = strings.Trim(line, " \t\r\n")
		if err != nil && len(line) == 0 { // End of File
			break
		}
		if len(line) == 0 || strings.IndexAny(line, "#") == 0 { //empty or comment
			continue
		}
		parts := strings.Split(line, ";")
		if len(parts) < 3 {
			return errors.New("invalid format")
		}
		if teams[parts[0]] == nil {
			teams[parts[0]] = &Score{team: parts[0]}
		}
		if teams[parts[1]] == nil {
			teams[parts[1]] = &Score{team: parts[1]}
		}
		action := actions[parts[2]]
		if action == nil {
			return errors.New("invalid action")
		}
		action(teams[parts[0]], teams[parts[1]])
	}

	scores := Scores{}
	for _, score := range teams {
		scores = append(scores, score)
	}
	sort.Sort(scores) //sort by points and name

	buffWriter := bufio.NewWriter(writer)
	buffWriter.WriteString(fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, score := range scores {
		buffWriter.WriteString(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n",
			score.team, score.matches(), score.wins, score.draws, score.looses, score.points()))
	}
	buffWriter.Flush()
	return nil
}
