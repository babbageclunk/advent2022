package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data := lib.Read("../input")
	lines := strings.Split(strings.TrimSpace(data), "\n")
	total := 0
	for _, group := range groups(lines) {
		common := findCommon(group)
		total += score(common)
	}
	fmt.Println(total)
}

func groups(lines []string) [][3]string {
	var groups [][3]string
	var group [3]string
	for i, line := range lines {
		group[i%3] = line
		if i%3 == 2 {
			groups = append(groups, group)
		}
	}
	return groups
}

func score(letter rune) int {
	if 'a' <= letter && letter <= 'z' {
		return int(letter-'a') + 1
	}
	return int(letter-'A') + 27
}

func runeSet(line string) lib.Set[rune] {
	result := lib.NewSet[rune]()
	for _, c := range line {
		result.Add(c)
	}
	return result
}

func findCommon(group [3]string) rune {
	s1 := runeSet(group[0])
	s2 := runeSet(group[1])
	s3 := runeSet(group[2])
	common := s1.Intersect(s2).Intersect(s3)
	if common.Len() != 1 {
		panic(fmt.Sprintf("expected 1 common letter, got %v", common))
	}
	return common.ToSlice()[0]
}
