package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data := lib.Read("input")
	lines := strings.Split(strings.TrimSpace(data), "\n")
	total := 0
	for _, line := range lines {
		mid := len(line) / 2
		c1 := line[:mid]
		c2 := line[mid:]
		common := findCommon([]rune(c1), []rune(c2))
		total += score(common)
	}
	fmt.Println(total)
}

func score(letter rune) int {
	if 'a' <= letter && letter <= 'z' {
		return int(letter-'a') + 1
	}
	return int(letter-'A') + 27
}

func findCommon(a, b []rune) rune {
	for _, letter := range a {
		for _, other := range b {
			if letter == other {
				return letter
			}
		}
	}
	panic(fmt.Sprintf("no common letter between %s and %s", string(a), string(b)))
}
