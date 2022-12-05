package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stacks := getStacks(Input)

	moves := getMoves(Input)
	for _, m := range moves {
		stacks = m.do(stacks)
	}

	for num, stack := range stacks {
		fmt.Println(fmt.Sprintf("%d: %s", num, stack))
	}

	numStacks := len(stacks)
	for i := 0; i < numStacks; i++ {
		stack := stacks[i]
		fmt.Println(stack[len(stack)-1])
	}
}

func getStacks(input string) map[int][]string {
	ret := map[int][]string{}

	parts := strings.Split(input, "\n\n")
	stackRows := strings.Split(parts[0], "\n")
	// remove the x axis label row and empty row at top
	stackRows = stackRows[1 : len(stackRows)-1]

	// Loop in reverse to order the stacks from bottom up
	for i := len(stackRows) - 1; i >= 0; i-- {
		row := stackRows[i]

		for j := 0; j < 9; j++ {
			k := 2 + (4 * j)
			if k >= len(row) {
				continue
			}
			char := row[k-1 : k]

			if char == " " {
				continue
			}

			stack, ok := ret[j]
			if !ok {
				stack = []string{}
			}
			stack = append(stack, string(char))
			ret[j] = stack
		}
	}

	return ret
}

type move struct {
	num  int
	from int
	to   int
}

func (m move) do(stacks map[int][]string) map[int][]string {
	from := m.from - 1
	to := m.to - 1

	stackFrom := stacks[from]
	stackTo := stacks[to]

	for i := 0; i < m.num; i++ {
		item := stackFrom[len(stackFrom)-1]
		stackFrom = stackFrom[:len(stackFrom)-1]
		stackTo = append(stackTo, item)
	}

	stacks[from] = stackFrom
	stacks[to] = stackTo

	return stacks
}

func getMoves(input string) []move {
	parts := strings.Split(input, "\n\n")
	movesRaw := strings.Split(parts[1], "\n")

	ret := []move{}

	for _, raw := range movesRaw {
		splits := strings.Split(raw, " ")
		num, _ := strconv.Atoi(splits[1])
		from, _ := strconv.Atoi(splits[3])
		to, _ := strconv.Atoi(splits[5])

		m := move{
			num:  num,
			from: from,
			to:   to,
		}

		ret = append(ret, m)
	}

	return ret
}
