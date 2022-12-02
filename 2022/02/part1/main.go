package main

import (
	"fmt"
	"strings"
)

// https://adventofcode.com/2022/day/2

const (
	PlayTypeScissors = "scissiors"
	PlayTypeRock     = "rock"
	PlayTypePaper    = "paper"
)

const (
	pointsWin  = 6
	pointsDraw = 3
)

type play struct {
	Type  string
	Score int
	Beats string
}

func (p play) pointsAgainstPlay(opponent play) int {
	switch opponent.Type {
	case p.Beats:
		return pointsWin
	case p.Type:
		return pointsDraw
	default:
		return 0
	}
}

var (
	playRock = play{
		Type:  PlayTypeRock,
		Score: 1,
		Beats: PlayTypeScissors,
	}
	playPaper = play{
		Type:  PlayTypePaper,
		Score: 2,
		Beats: PlayTypeRock,
	}
	playScissors = play{
		Type:  PlayTypeScissors,
		Score: 3,
		Beats: PlayTypePaper,
	}
)

func getPlay(s string) (play, error) {
	switch s {
	case "A", "a", "X", "x":
		return playRock, nil
	case "B", "b", "Y", "y":
		return playPaper, nil
	case "C", "c", "Z", "z":
		return playScissors, nil
	}

	return play{}, fmt.Errorf("Unknown play: %s", s)
}

func main() {
	score, err := run()
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("Player 2 total score: %d", score)
}

func run() (int, error) {
	points := 0

	rounds := strings.Split(Input, "\n")

	for _, round := range rounds {
		plays := strings.Split(round, " ")
		if len(plays) != 2 {
			return points, fmt.Errorf("Invalid round: %s", round)
		}
		player1, err := getPlay(plays[0])
		if err != nil {
			return points, fmt.Errorf("Invalid play: %s", plays[0])
		}

		player2, err := getPlay(plays[1])
		if err != nil {
			return points, fmt.Errorf("Invalid play: %s", plays[1])
		}

		points += player2.Score
		points += player2.pointsAgainstPlay(player1)
	}

	return points, nil
}
