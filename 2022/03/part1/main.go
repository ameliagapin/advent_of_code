package main

import (
	"fmt"
	"strings"
)

func main() {
	splits := strings.Split(Input, "\n")

	bothSidesSum := 0

	rucksacks := []rucksack{}
	for _, line := range splits {
		rs := getRucksack(line)
		rucksacks = append(rucksacks, rs)

		both := rs.getItemsInBothSides()

		for b, _ := range both {
			bothSidesSum += getPriority(b)
		}
	}

	fmt.Println("Sum: ", bothSidesSum)
}

type rucksack struct {
	sideOne map[rune]int
	sideTwo map[rune]int
}

func (r rucksack) getItemsInBothSides() map[rune]int {
	ret := make(map[rune]int)

	for item, numS1 := range r.sideOne {
		if numS2, ok := r.sideTwo[item]; ok {
			ret[item] = numS1 + numS2
		}
	}

	return ret
}

func getRucksack(items string) rucksack {
	ret := rucksack{
		sideOne: make(map[rune]int),
		sideTwo: make(map[rune]int),
	}

	num := len(items)
	for _, s := range items[:num/2] {
		num, _ := ret.sideOne[s]
		num++
		ret.sideOne[s] = num

	}
	for _, s := range items[num/2:] {
		num, _ := ret.sideTwo[s]
		num++
		ret.sideTwo[s] = num
	}

	return ret
}

func getPriority(item rune) int {
	ascii := int(item)
	if ascii > 96 {
		return ascii - 96
	}

	return ascii - 64 + 26
}
