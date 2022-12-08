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

	visibleTrees := make(map[string]int64)
	for y, row := range forest {
		for x, height := range row {
			if forest.isVisible(x, y) {
				visibleTrees[getTreeKey(x, y)] = height

				fmt.Println(fmt.Sprintf("%d | %d : visible", y, x))
			} else {
				fmt.Println(fmt.Sprintf("%d | %d : not visible", y, x))
			}
		}
	}

	fmt.Println(fmt.Sprintf("Visible trees: %d", len(visibleTrees)))
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

func (f forest) isVisible(x, y int) bool {
	if x == 0 || y == 0 {
		return true
	}
	if x == len(f[y])-1 || y == len(f)-1 {
		return true
	}
	return f.isVisibleNorth(x, y) || f.isVisibleSouth(x, y) || f.isVisibleWest(x, y) || f.isVisibleEast(x, y)

}

func (f forest) isVisibleNorth(x, y int) bool {
	treeHeight := f[y][x]

	for i := (y - 1); i >= 0; i-- {
		h := f[i][x]
		if h < treeHeight {
			continue
		}

		return false
	}

	return true
}

func (f forest) isVisibleSouth(x, y int) bool {
	treeHeight := f[y][x]

	for i := (y + 1); i < len(f); i++ {
		h := f[i][x]
		if h < treeHeight {
			continue
		}
		return false
	}

	return true
}

func (f forest) isVisibleWest(x, y int) bool {
	treeHeight := f[y][x]

	for i := (x - 1); i >= 0; i-- {
		h := f[y][i]
		if h < treeHeight {
			continue
		}
		return false
	}

	return true
}

func (f forest) isVisibleEast(x, y int) bool {
	treeHeight := f[y][x]

	for i := (x + 1); i < len(f[y]); i++ {
		h := f[y][i]
		if h < treeHeight {
			continue
		}
		return false
	}

	return true
}

func getTreeKey(x, y int) string {
	return fmt.Sprintf("%6d|%6d", y, x)
}
