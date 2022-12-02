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

	rounds := strings.Split(in, "\n")

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

var in = `A Y
B X
B X
C Y
B X
A Z
B X
B X
C Z
A Z
C Z
A X
A Y
A Y
B X
C Y
B Z
A X
B X
C Y
C Y
A Y
A Y
C Y
B X
B X
B X
A X
B X
B X
C Y
C Y
B X
A Z
A Z
B X
A Z
C Y
B X
B X
B X
B X
B X
B Z
A Z
B X
A Z
A X
B X
A Z
B X
C Y
A Z
A Z
B X
B X
B Z
C Y
A Y
B Z
A Z
C Z
B X
C Y
C Y
B Z
B X
B X
B Z
B X
B X
A Y
A Z
B Z
B X
B Y
B Z
B X
C X
B Z
B X
A Z
B Z
B Z
B X
A Y
A Z
A Y
B X
B Z
B X
B X
A Y
A Y
B X
B X
B X
C X
C Y
B X
B X
B Z
B X
B X
B X
C Y
B X
C Y
A Z
B X
B Z
B X
C Y
C Y
A Z
A Y
C Y
B X
B X
B Z
A Y
B X
B X
B X
A Y
B Z
B X
A Y
B Z
C Y
A X
B X
B X
A Z
C Y
B X
B Z
A Y
A Y
C Y
A Y
C Y
B X
B Z
B X
C Y
A Y
B X
C Z
A Z
C Y
B X
B Z
B X
B X
B Z
B Y
A Y
B X
A Y
A Z
A Z
B X
C Y
B Z
A X
A Y
B X
A Z
B X
A Y
B Y
B X
A Y
B X
B X
B X
B X
B X
B X
B X
A Z
B X
B X
B Z
B Z
B Z
B Z
A Y
B Z
C Z
B X
B X
B X
A Z
B Z
C Y
B X
C Y
B Z
B X
B Z
B X
C Z
A Z
A Y
A X
B X
B Z
A Y
B X
B X
B X
B X
C Y
C Y
A Z
A Y
A X
B X
B X
A Y
B X
C Y
A X
B X
B X
B X
A Z
C Y
B Z
C Y
A Y
C Y
A X
A Z
C Z
A Z
B X
B Z
B X
B X
A X
B X
B X
B Z
A Y
C Y
A Z
C Z
A Z
A Z
A Z
B Z
B Z
B Z
C Z
B Y
C Y
C Y
C X
B X
C Y
B X
B X
C Y
B Z
B X
C Z
B X
A Z
B Z
C X
C Y
B X
A Y
A Z
B X
A Z
A Z
B X
B X
C Z
B Z
A Z
A Z
C Z
B X
A Z
A X
B X
B X
C Y
B X
B Z
A Z
B X
B X
B X
B X
B X
B X
B X
B X
C Y
C Y
B X
B X
B X
C Y
A X
B Z
B X
B X
C Z
B X
B Z
B Z
A Z
B X
B Z
B X
A Z
B X
B X
B X
B X
B X
C Z
B X
B X
B X
B Z
C Y
C Y
A Y
B X
B Z
C Y
B X
B X
C Y
C Y
B X
B X
B X
B X
C X
B Z
B X
B Z
C Y
A X
A Z
B X
B X
A Z
B X
C Z
B X
A Y
A Z
B X
B X
B X
B Z
B X
A Y
C Y
B Z
B Z
C Y
B X
A Y
C Y
B X
C Z
B X
B X
A Z
B Z
B Z
B X
A Z
A Z
A Y
B Z
B X
B X
B X
B X
B Z
B X
B Z
B X
A Z
B X
B X
C Z
C Z
B X
C Y
A Z
B X
B Z
C Z
C X
B X
A Y
A Y
B X
B X
C Y
A X
C Y
A Y
A Z
C Z
B Z
A Z
B X
B X
B Z
C Y
A Y
A Y
B Z
A Z
A Z
B Z
A Z
B X
A Y
A Y
C Y
B X
B X
A X
B X
C X
A Y
A Z
B X
B X
B X
B X
B X
B Z
B Z
B Z
B X
B X
A Z
B X
C Z
A Z
C Y
B Z
C Y
A Y
C Y
B X
B X
B X
A Y
B Z
B Z
B X
B Z
B X
B X
B X
C Y
B X
B Z
A Z
B X
B X
B X
C Z
C Z
B Z
B X
A Z
A X
B X
B X
B Y
A Y
B X
B X
C Z
A Z
B X
B Z
B Z
A Y
A Z
B Z
B Z
A Y
A Y
A Z
C Y
C Y
B X
B X
A Z
B X
C Y
C X
A Z
B X
B X
B X
A Y
C Y
B Y
C Y
C X
A X
A Y
B Z
B X
B Z
A X
B X
C Z
B X
B X
A Y
B Z
B X
A Y
A Y
B X
C Y
B X
B X
C Y
B X
C Z
B Z
B X
A X
C Y
C Y
B X
B Z
B X
B X
A Z
B X
A Z
B Z
B X
B X
B Z
A Y
C Y
B Z
B X
A Y
B X
A X
C Z
B X
B Z
B Z
A X
C Y
A Z
B X
B Z
A Y
B X
C Y
A Y
B Z
C Y
B X
B X
B X
B X
B X
B X
C Y
B X
C Z
A Y
C Y
B X
C Y
B Z
B X
A X
B Z
B X
C X
B X
A Z
B X
A Y
B X
C Z
B X
B Z
A Y
B X
A Z
A Z
A Z
A Y
A Y
B X
A Z
B X
B X
A Y
A Y
A Y
B Z
C Y
B Z
B X
C Z
A Y
C Y
B X
B X
B Z
B X
B X
B X
C Y
C Y
A Y
B Z
B Z
B X
B X
B Z
B X
A Y
C Y
B X
C X
B Z
B X
B Z
B Z
C Y
B Z
B X
C X
B X
C Y
B X
A Z
B X
B X
B X
B X
B X
C Y
A Y
B X
C Y
B Z
B Z
B X
C X
B Z
C Y
C X
C Y
A Z
B Z
C Y
C X
B X
A Y
C Y
B X
C Z
B Z
A Y
C Z
C X
B X
B X
A Y
C X
B X
B X
A X
A Z
A X
A Y
A Y
C Y
B X
C Z
B X
C Y
A Y
B X
B X
B X
B X
A Z
C X
B X
B Z
B X
B X
C Y
B X
C Z
B X
B X
B X
B X
B X
C Y
B X
B X
A Y
B X
C Y
C Y
A Y
B X
A X
A Y
B X
B Z
B X
B Z
A X
A Y
A X
B X
C Z
A Y
B Z
A Y
B Z
B Z
B X
B Z
B X
B X
B X
B Z
B X
B X
B X
A Z
C Z
C Z
A X
A X
B X
B X
A X
B X
A X
A Z
B Z
B X
C Y
B X
C Y
A Y
C X
A X
A Y
B Z
B X
B X
C Z
B Y
A Y
B Y
C Y
A Y
B X
B X
C X
B Z
A Y
B Z
A Z
B X
B X
A Z
C Z
A Y
B X
B X
B Z
B X
B X
B Z
A X
A Z
C Y
A Y
A Z
B X
C Y
B Z
C Y
B X
B X
B Z
B X
A Y
B X
A Y
B X
A Y
B X
A Y
C Y
A Y
B Z
A Y
B X
A Z
C Y
B X
B X
B X
A X
B X
B Z
A Z
A Z
B X
B Z
B X
A Y
B Z
B X
B X
A Y
A Y
C Z
C Y
B X
A X
A Y
A X
C Y
B X
C X
B X
B X
A Y
A Y
A Y
B Y
B X
B X
B X
C Z
B Z
C Z
A X
B X
B X
A Y
A Z
B X
A Y
B Z
B Z
C Y
B X
B X
C X
B X
C X
B X
B Z
A X
B X
A Z
A X
B Z
A Y
B X
B X
B X
B X
B Z
A Z
B X
B Z
B X
B Z
C Y
B X
C Y
B X
B X
B X
A Y
B X
A Y
B X
B X
A X
C Z
C Y
A Z
B X
A Y
C Z
C Y
C Y
B Z
A Y
B X
A Y
B X
B X
B Z
B Z
B X
C X
B Z
A X
A Y
A Y
B Z
A Z
C Y
B X
B Z
A Y
C Y
A Y
A X
A Z
B X
B X
C Y
B X
C Y
B X
A X
B Z
C X
B Y
A Y
C Y
B X
A X
B X
B X
B X
A Z
A Y
B X
C Z
B X
C X
C Z
B X
B X
C Y
B X
B Z
A Y
A Y
B X
B X
B Z
C Y
B X
B X
B X
A Y
C Y
B X
A Y
A Z
C Y
B X
A Y
A Z
B Z
A Z
B X
C X
B Z
A Y
B Z
A Y
C Y
A X
A Y
B X
C Y
C Y
B X
A X
B Z
B X
B X
A Z
A Z
A Y
A Y
B X
B X
B X
A X
B X
B X
B X
A X
B Z
C Z
A Y
C Y
C Y
C X
B X
B Z
B Z
B X
B X
C Y
A X
C Y
A X
B X
B X
B Z
B Z
C Y
C Y
A Z
C Y
A Z
B X
A Z
A Y
B Z
B Z
B X
B X
A Y
B Z
B X
C Z
A Z
B Z
B X
B Z
B X
C Y
B X
A Y
B Z
C Y
B Z
C Z
B X
C Z
B X
B X
B X
C Y
B X
B X
B Z
A Y
B X
B Z
A Z
B X
C Y
C X
C Y
A Z
B X
A Z
B X
C Z
A Y
B Z
B Z
A Y
C Z
C X
A Y
B X
B Z
B X
B X
B X
B X
A X
B X
B X
A Z
B X
B X
A Z
C Y
A X
B Z
B X
B X
A Z
B X
A Y
B Z
C Y
B X
A Y
C Y
B Z
B X
A Z
A Z
B X
B X
B Z
B X
B X
A Z
B X
B Z
B Z
C Y
A Z
B X
A Y
B X
B X
B X
B Z
A X
C Y
B Z
B X
B X
B X
B X
B X
B Z
B X
B Z
A Z
B Z
B X
C Z
B Z
C Z
A Z
C X
A Y
B Z
B Z
B X
B Z
B Z
A Z
C Z
A X
C Y
C Y
A Y
A X
A Z
B X
A Y
B X
B X
C Y
C Y
B X
B Z
A Y
B X
B Z
B X
B X
A Z
C Y
B X
C Y
A Y
B Z
A Z
A Y
B X
A Z
B Z
B X
A Y
A Y
A Z
A X
B Z
C Z
B X
C Z
B X
B Z
A X
B X
B X
B X
B X
C Y
B X
A X
B Z
A Y
B X
A Z
B X
B Z
A Y
A Y
A Y
A Y
A Z
C Y
C Z
B X
B Z
A Y
B X
B X
B X
A Y
B Z
B X
B Z
B X
B X
B X
C Y
C Y
C X
B X
A Y
A Y
C Z
A Y
A Y
C X
A Z
B X
B X
B X
B X
C Y
A Y
B Z
A Z
B X
C X
A Z
B X
C Y
B X
B X
C Y
B X
B X
B X
A Y
B X
B X
C Z
B X
C Y
C Y
B X
A Z
B X
B X
C Y
A Y
B X
B Z
B X
B X
B X
C Y
C Z
B Z
C Y
B X
B Z
B X
B X
B X
C Y
C X
B X
A Y
A Z
B Z
A Y
B X
B X
B Y
C Y
C X
B Z
A Z
B X
B X
A Y
C Z
C Y
B X
C Y
A Z
A Y
B Z
B X
B X
B X
B X
A Z
A X
A X
A X
C Y
B Z
B X
C Y
B X
B X
A Z
C Z
C Z
B X
A Y
B X
B Z
B X
A Y
A Y
A Z
C Y
A Y
A Z
A Y
A Y
B X
A Y
B X
A Z
A Z
C Y
B X
B X
C Y
B X
B X
B X
C Y
A Y
B X
A Y
B Z
C Y
B X
A Y
B Z
A Z
B Z
A Z
B X
C Y
B X
B X
B X
C Y
A Y
B X
A Y
C X
A Z
B X
B Y
B X
B X
A Z
B X
B X
B Z
B X
C Y
B X
B X
B Z
B X
C Y
B X
C Y
B X
B Z
A Y
B X
B X
C Z
B X
C Y
A Y
C Y
B X
B Z
B X
B X
B X
B X
A Y
B X
A Z
A Y
A X
C Y
B X
C Y
B X
A Z
C Y
C Y
B Z
B X
A Y
B Z
A Y
A Y
B Z
B Z
B X
C Y
C Y
B X
A X
A Z
B X
C Y
B X
B X
B Z
B X
B Z
C Y
C Z
C Y
A Y
A Z
B X
C Y
A Y
C X
B X
B X
A Y
B X
B X
B X
B Z
B Y
A X
A Y
C Y
C Y
A Y
C Y
B Z
B Z
C Y
B X
C Y
C Z
A Y
C X
A Y
A Y
C Y
C Y
B X
A Y
A X
B X
B X
B Z
B X
A Z
A Y
A Y
B X
B X
B Z
B Z
B X
B Z
A Z
A Y
A Z
B X
C Y
B Z
B X
A Y
B X
C X
B X
B X
C Y
B X
B X
B X
A Y
A X
A Y
B X
B X
A Y
A Y
A Z
A Z
C X
C Y
A Z
C Y
B Z
B X
A Y
B X
B Z
C Z
A Y
A Z
A Y
B X
B Z
B X
A X
B Z
B X
C Y
B X
A Z
C Z
A X
B X
B X
A Y
C Y
B X
B X
B X
B X
B X
A X
B X
A Z
C Z
B Z
B X
C X
A X
C Y
A Y
A Z
B X
A X
A Z
A Z
B X
B X
A Z
C Y
A Z
B X
C Y
B X
C Z
C Y
C X
A Z
C X
B Z
B X
B X
C Y
B X
C Y
A Z
A Z
A Z
A Y
A Z
B X
B X
B X
B Z
C Y
C Z
C Y
C Y
C Y
A Y
B Z
A X
B X
B X
B X
A X
A Y
B X
B Z
A Z
A Y
A Y
B Z
B X
B X
C Y
C Y
C Y
B X
C Y
B X
B X
B X
C Y
A Y
A Y
C Z
C Y
B X
A X
A Y
B X
C Y
B X
A Y
C Y
A Y
B X
B X
B Z
A Y
B Z
B Z
B X
B X
C X
C Y
B X
B X
A X
C Y
B X
B Z
B X
B X
A Y
B X
B Z
B Z
C Y
B X
B Z
B Z
B Z
B X
B X
A Y
B X
B X
B Z
C Y
A Y
B X
B Z
B X
B X
A X
B X
B X
B X
B Z
B Z
B Z
B X
A Z
B X
A Z
A Y
B Z
B X
A Y
B X
B Z
B X
A Z
B X
B Z
C Z
A X
B X
B X
B X
C Y
B X
B X
A Y
A Z
B X
C Y
B X
C Y
B X
B Y
B X
C Z
B X
C Z
A Y
B X
C Z
B X
A Y
C Y
B X
B X
B Z
C X
B Z
B X
A Z
B X
B X
B X
B Z
B X
A Z
C Y
A Y
B X
B X
A Z
B X
A Y
B X
C Y
B Z
B Z
B X
B Z
B X
B X
B X
B X
B X
C Y
A Z
A Z
B Z
A Y
B Z
C Y
A Y
A Y
C Y
A Y
B Z
A Y
A Y
B X
B X
A X
B X
C Y
C Y
B Z
B X
A Z
B Z
B X
A Z
B Z
A Y
A Y
A Y
A Y
A Z
A Y
B X
B Z
A Y
B X
A Z
B Z
C X
B X
B X
B Z
B Z
B X
B X
B X
B X
B X
B X
B X
B Z
B X
B X
B X
B X
B X
B Z
B X
B X
B Z
B Z
B X
A Y
B Z
B X
A Z
A Y
A Z
B X
B X
A X
B Y
A Z
B X
C Z
A Z
B X
B Z
B X
B X
B X
A Y
C Y
B X
B X
B X
A Y
B X
A Z
B X
B X
B X
B Z
B Z
C X
C Y
B X
B X
C Y
B X
B X
B X
A X
A Z
A Z
B Z
B Z
B X
A Y
B X
C Y
B Z
C Y
C Z
A Y
A Y
B X
B X
A Z
B X
B X
B Z
B X
B X
A Z
B X
B X
B X
C Y
B X
B X
C Y
B X
B X
B Z
B X
B X
A Y
C Y
A X
C Y
B X
B X
A Z
A Y
B X
A X
B Z
A Y
B X
B X
A Y
B X
C X
B X
B X
B Z
B Z
B X
B Z
B X
B X
A Y
A Y
B X
B Z
C Y
B X
B X
B X
A Z
B X
A Y
B X
C Y
A X
A Y
B X
B X
C Y
B X
B X
B X
A Z
B Z
B X
C X
A Z
C Y
B X
B Z
B X
B X
A Z
B X
A Y
C Y
B X
B Z
C X
A Z
A Z
B Z
B Z
A X
B X
A Y
B X
B X
B X
B X
A Z
B Z
B Z
C Y
A X
A Y
A Y
A Z
B X
A Y
B Z
C Y
B Z
B X
A Z
B Z
B X
B Z
B Z
B Z
B X
C Y
B X
B X
A Y
B X
C X
A Y
A Z
A Y
B X
C Y
B Z
B X
B X
A X
B X
B X
A Y
A Z
C Z
A Y
B X
B X
C Y
A X
A Z
B X
B Z
B Z
C Y
A Z
C Y
A Z
C Y
C Y
B Z
C Y
B X
C Y
C Y
B X
B Z
C X
C Y
A Y
B X
C Y
B X
B X
C Y
B Z
B Z
A Z
A Z
C Y
B X
B Z
B X
B X
B X
C Y
B X
C Y
A Z
A Z
B X
B X
A Y
A X
B X
A Z
A Y
C Y
B X
A Y
A Y
B X
A X
A Z
A Y
A Z
C Y
C Y
C Y
A Y
B X
B X
A Y
C Y
B X
A X
C Y
C Y
A Y
C Y
C Y
B Z
C Y
B X
C Y
B X
B X
B X
A Y
B Y
B X
C X
B X
A X
A Y
C Y
C Y
B X
B X
A Y
B X
B X
A Y
C Y
C X
A Y
B X
B Z
B Z
B Z
B X
A Z
B X
B X
A Y
B X
B X
A Y
B Z
A Y
B X
B X
B Z
A Z
A Y
A Z
A Y
C Y
B X
C X
B X
B X
A Z
B X
B Z
B X
A Z
C Y
B X
B X
B Z
A Z
A Y
B X
B Y
A Y
A Z
A Y
C Z
B X
B X
B X
A Z
B X
B X
B X
A Y
B Z
C Y
C Y
A Z
A Z
A Y
B X
A Y
B Z
C Y
B X
B Z
A Y
B X
A Y
C Y
B X
B X
A Y
A Z
A Z
A Z
B X
B X
A Z
B X
B X
B Z
A X
B Z
A Z
B X
C Y
C X
A Z
B X
A Y
B Y
B Z
B X
B X
A Y
B X
A Z
A Z
C X
B X
B Z
A Y
C Y
A Y
B X
B Z
B X
A Z
B Z
C Y
B Z
B X
B X
B X
A Y
B X
B X
A Z
C Y
B X
C Y
B X
C Y
A Z
A Y
A X
C Y
B X
B X
B Z
B X
A Y
A Y
B X
B X
A Z
C Y
B Z
B X
B X
A X
A Y
B Z
C Y
A Z
C Y
A Z
B X
A Y
B X
A Y
B X
B X
C X
B Y
C Z
A Z
C Y
B X
B X
B Z
C Z
A X
A Y
A Z
B X
C Y
B Z
B Z
A Y
B X
B X
A Y
A Z
B X
B X
A Z
B X
B X
B Z
A X
B X
A X
A Y
C Y
B X
B Z
A Y
B X
A X
B X
B X
B Z
B X
B X
A Y
B X
B X
A Y
A Y
C Y
B Z
B X
A X
A Z
B X
B X
C Z
B X
B X
B X
B X
B X
C Y
C Y
C X
B X
B X
B Z
B Z
B X
A Y
B Z
B Z
B Z
B X
B X
B X
B X
A Y
A Y
B X
C Y
C Y
B X
B X
A Y
C Y
A Z
C Y
B X
C Z
B X
B Z
B X
C Y
C Y
B X
B X
B X
C Y
B X
B X
A X
B X
C Y
A Y
B X
B X
B X
C X
C Y
A Y
B X
B X
C Y
A Y
B X
A Y
A Y
B Z
C Y
B Z
A Y
A Z
C Y
B X
C Y
B X
C Y
C Y
B X
B X
B X
B Z
A Y
B X
B X
B X
A Y
B X
B X
A Y
B X
A Y
C Y
A Z
B X
B X
C Y
B X
B X
A Y
A Z
C Y
B X
B Z
B Z
B X
C Y
B X
A Y
C X
C Y
B X
B X
B Z
C Y
B X
B X
C Y
C Y
B X
C Y
B Z
C Y
C X
B Z
C Y
B Z
C Y
A X
B Z
B Z
B X
B X
C Y
B X
C Y
B X
B X
B X
A Y
B X
C Y
B Z
C Y
C Y
A Y
C Z
A Y
B X
A Z
C Z
C Y
B X
B X
A Y
B Z
B X
C Y
B X
B Z
A Z
B Z
A Y
C Y
A Z
B X
B X
C Y
C Y
B X
B X`
