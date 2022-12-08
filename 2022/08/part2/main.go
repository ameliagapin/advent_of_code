package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	forest := make(forest)

	rows := strings.Split(Input, "\n")
	for y, row := range rows {
		for x, height := range row {
			forest = forest.addTree(x, y, string(height))
		}
	}

	scores := make(map[string]int64)
	for y, row := range forest {
		for x, _ := range row {
			scores[getTreeKey(x, y)] = forest.getScore(x, y)
			fmt.Println(fmt.Sprintf("%d | %d : %d", y, x, scores[getTreeKey(x, y)]))
		}
	}

	highest := int64(0)
	for _, v := range scores {
		if v > highest {
			highest = v
		}
	}

	fmt.Println(fmt.Sprintf("Highest Score: %d", highest))
}

type forest map[int]map[int]int64

func (f forest) addTree(x, y int, height string) forest {
	h, _ := strconv.ParseInt(height, 10, 64)

	if _, ok := f[y]; !ok {
		f[y] = make(map[int]int64)
	}
	f[y][x] = h

	return f
}

func (f forest) getScore(x, y int) int64 {
	return f.getScoreNorth(x, y) * f.getScoreSouth(x, y) * f.getScoreWest(x, y) * f.getScoreEast(x, y)

}

func (f forest) getScoreNorth(x, y int) int64 {
	score := int64(0)
	treeHeight := f[y][x]

	for i := (y - 1); i >= 0; i-- {
		score++

		h := f[i][x]
		if h < treeHeight {
			continue
		}

		return score
	}

	return score
}

func (f forest) getScoreSouth(x, y int) int64 {
	score := int64(0)
	treeHeight := f[y][x]

	for i := (y + 1); i < len(f); i++ {
		score++

		h := f[i][x]
		if h < treeHeight {
			continue
		}
		return score
	}

	return score
}

func (f forest) getScoreWest(x, y int) int64 {
	score := int64(0)
	treeHeight := f[y][x]

	for i := (x - 1); i >= 0; i-- {
		score++

		h := f[y][i]
		if h < treeHeight {
			continue
		}
		return score
	}

	return score
}

func (f forest) getScoreEast(x, y int) int64 {
	score := int64(0)
	treeHeight := f[y][x]

	for i := (x + 1); i < len(f[y]); i++ {
		score++

		h := f[y][i]
		if h < treeHeight {
			continue
		}
		return score
	}

	return score
}

func getTreeKey(x, y int) string {
	return fmt.Sprintf("%6d|%6d", y, x)
}
