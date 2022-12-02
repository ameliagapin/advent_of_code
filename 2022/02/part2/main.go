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

const (
	outcomeWin  = "win"
	outcomeLose = "lose"
	outcomeDraw = "draw"
)

type play struct {
	Type  string
	Score int
	Beats string
}

func (p play) losesTo() string {
	for _, op := range []play{playRock, playPaper, playScissors} {
		switch op.Type {
		case p.Type:
			continue
		case p.Beats:
			continue
		default:
			return op.Type
		}
	}
	return ""
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

func (p play) getPlayForOutcome(outcome string) play {
	playType := ""
	switch outcome {
	case outcomeWin:
		playType = p.losesTo()
	case outcomeLose:
		playType = p.Beats
	default:
		return p
	}

	// this can't error because we know what we're passing it
	ret, _ := getPlay(playType)
	return ret
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
	case "A", "a", PlayTypeRock:
		return playRock, nil
	case "B", "b", PlayTypePaper:
		return playPaper, nil
	case "C", "c", PlayTypeScissors:
		return playScissors, nil
	}

	return play{}, fmt.Errorf("Unknown play: %s", s)
}

func getOutcome(s string) (string, error) {
	switch s {
	case "X", "x":
		return outcomeLose, nil
	case "Y", "y":
		return outcomeDraw, nil
	case "Z", "z":
		return outcomeWin, nil
	}
	return "", fmt.Errorf("Unknown outcome: %s", s)
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

		outcome, err := getOutcome(plays[1])
		if err != nil {
			return points, fmt.Errorf("Invalid outcome: %s", plays[1])
		}

		player2 := player1.getPlayForOutcome(outcome)

		points += player2.Score
		points += player2.pointsAgainstPlay(player1)
	}

	return points, nil
}
