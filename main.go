package main

import (
	"fmt"
	"os"

	"github.com/PjHiK7t5J/task-onefootball/processor"
	"github.com/PjHiK7t5J/task-onefootball/vintagemonster"
)

func main() {
	// initialize list of teams to find
	teamsToFind := map[string]bool{
		"Germany":          true,
		"England":          true,
		"France":           true,
		"Spain":            true,
		"Manchester Utd":   true,
		"Arsenal":          true,
		"Chelsea":          true,
		"Barcelona":        true,
		"Real Madrid":      true,
		"FC Bayern Munich": true,
	}

	// initialize vintagemonster REST client
	vmonster := vintagemonster.New()

	// initialize processor
	proc := processor.New(vmonster)

	// process requested teams
	listOfPlayers, err := proc.Process(teamsToFind)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// print result to the STDOUT
	proc.Print(listOfPlayers)
}
