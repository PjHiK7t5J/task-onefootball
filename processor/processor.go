package processor

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/PjHiK7t5J/task-onefootball/player"
	"github.com/PjHiK7t5J/task-onefootball/vintagemonster"
)

// Processor is an engine. It's main job is to work with players.
// It knows how to fetch them and how to print them.
type Processor struct {
	vmonster *vintagemonster.Vintagemonster
	writer   io.Writer
}

// New return new Processor
func New(vmonster *vintagemonster.Vintagemonster) *Processor {
	return &Processor{
		vmonster: vmonster,
		writer:   os.Stdout,
	}
}

// Process do the main logic of fetching players by requested teams
func (p *Processor) Process(teamsToFind map[string]bool) (player.PlayersList, error) {
	// initialize map to save all players
	allPlayers := map[string]*player.Player{}

	// initialize counters
	teamsCount := len(teamsToFind)
	teamsFoundCount := 0
	teamID := 0

	// search for teams and collect players
	for teamsCount != teamsFoundCount {
		teamID++

		res, err := p.vmonster.Get(teamID)
		if err != nil {
			return nil, err
		}

		// check, that this is a team that we are searching for.
		// If not, skip all the following steps.
		_, ok := teamsToFind[res.Data.Team.Name]
		if !ok {
			continue
		}

		teamsFoundCount++

		for _, p := range res.Data.Team.Players {
			_, ok := allPlayers[p.Name]
			if !ok {
				player := &player.Player{}
				player.Name = p.Name
				player.Age = p.Age

				allPlayers[p.Name] = player
			}
			allPlayers[p.Name].AddTeamUnique(res.Data.Team.Name)
		}
	}

	players := player.PlayersList{}
	for _, v := range allPlayers {
		players = append(players, v)
	}

	sort.Sort(players)
	return players, nil
}

// Print prints list of of players in specified format:
// full name; age; list of teams
func (p *Processor) Print(listOfPlayers player.PlayersList) {
	for i, v := range listOfPlayers {
		fmt.Fprintf(p.writer, "%d. %s; %s; %s\n", i+1, v.Name, v.Age, strings.Join(v.Teams, ", "))
	}
}
