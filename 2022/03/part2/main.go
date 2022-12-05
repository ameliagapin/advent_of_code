package main

import (
	"fmt"
	"strings"
)

func main() {
	splits := strings.Split(Input, "\n")

	rucksacks := []rucksack{}
	for _, line := range splits {
		rs := getRucksack(line)
		rucksacks = append(rucksacks, rs)
	}

	total := 0

	for i := 0; i < len(rucksacks); i++ {
		rs := rucksacks[i]
		i++
		rs2 := rucksacks[i]
		i++
		rs3 := rucksacks[i]

		var badge rune

		for item, _ := range rs.items {
			if !rs2.hasItem(item) {
				continue
			}
			if !rs3.hasItem(item) {
				continue
			}
			badge = item
		}

		total += getPriority(badge)
	}

	fmt.Println("Sum: ", total)
}

type rucksack struct {
	items map[rune]int
}

func (r rucksack) hasItem(item rune) bool {
	_, ok := r.items[item]
	return ok
}

func getRucksack(items string) rucksack {
	ret := rucksack{
		items: make(map[rune]int),
	}

	for _, s := range items {
		num, _ := ret.items[s]
		num++
		ret.items[s] = num
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
