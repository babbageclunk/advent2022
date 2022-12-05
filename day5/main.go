package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	stacks, rest := readCrates(lines)
	moves := readMoves(rest)
	for _, m := range moves {
		m.apply(stacks)
	}
	for _, stack := range stacks {
		fmt.Printf("%s", stack[len(stack)-1])
	}
	fmt.Println()
}

type stacks [][]string

type move struct {
	n, from, to int
}

func pop(items []string) (string, []string) {
	return items[len(items)-1], items[:len(items)-1]
}

func push(items []string, item string) []string {
	return append(items, item)
}

func (m move) apply(s stacks) {
	var item string
	for i := 0; i < m.n; i++ {
		item, s[m.from] = pop(s[m.from])
		s[m.to] = push(s[m.to], item)
	}
}

func readCrates(lines []string) (stacks, []string) {
	// luckily the lines have trailing whitespace
	numStacks := (len(lines[0]) + 1) / 4
	result := make(stacks, numStacks)
	last := 0
	for i, line := range lines {
		if line == "" {
			last = i
			break
		}
		for j := 0; j < numStacks; j++ {
			if line[j*4] == ' ' {
				continue
			}
			charPos := (j * 4) + 1
			result[j] = append(result[j], line[charPos:charPos+1])
		}
	}

	// reverse the stacks so the bottom is the first element
	for _, stack := range result {
		reverse(stack)
	}

	return result, lines[last+1:]
}

func reverse(items []string) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func readMoves(lines []string) []move {
	result := make([]move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		result[i] = move{
			n:    lib.ParseInt(parts[1]),
			from: lib.ParseInt(parts[3]) - 1,
			to:   lib.ParseInt(parts[5]) - 1,
		}
	}
	return result
}
