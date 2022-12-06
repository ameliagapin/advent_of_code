package main

import "fmt"

func main() {

outter:
	for i, _ := range Input {
		if i < 14 {
			continue
		}

		m := make(map[rune]bool)
		for _, p := range Input[i-14 : i] {
			if _, ok := m[p]; ok {
				continue outter
			}
			m[p] = true
		}
		fmt.Printf("Index: %d", i)
		break
	}
}
