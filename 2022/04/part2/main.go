package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	groups := strings.Split(Input, "\n")

	overlapping := 0
	for _, group := range groups {
		elf1, elf2 := getAssigments(group)
		if isOverlapping(elf1, elf2) {
			overlapping++
		}
	}

	fmt.Println("Fully contained: ", overlapping)
}

func getAssigments(in string) ([]string, []string) {
	split := strings.Split(in, ",")

	elf1 := strings.Split(split[0], "-")
	elf2 := strings.Split(split[1], "-")

	return elf1, elf2
}

func isOverlapping(elf1, elf2 []string) bool {
	elf1Start, _ := strconv.ParseInt(elf1[0], 10, 64)
	elf1End, _ := strconv.ParseInt(elf1[1], 10, 64)
	elf2Start, _ := strconv.ParseInt(elf2[0], 10, 64)
	elf2End, _ := strconv.ParseInt(elf2[1], 10, 64)

	if elf1Start <= elf2Start && elf1End >= elf2End {
		return true
	}
	if elf2Start <= elf1Start && elf2End >= elf1End {
		return true
	}
	if elf1Start >= elf2Start && elf1Start <= elf2End {
		return true
	}
	if elf1End >= elf2Start && elf1End <= elf2End {
		return true
	}

	return false
}
