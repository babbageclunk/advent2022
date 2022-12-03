package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data, err := ioutil.ReadFile("input")
	lib.Check(err)
	lines := strings.Split(string(data), "\n")
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		total += Round{
			opponent: parts[0],
			you:      parts[1],
		}.score()
	}
	fmt.Println(total)
}

// Rock - A X
// Paper - B Y
// Scissors - C Z
var beats map[string]string = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var points map[string]int = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var same map[string]string = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

type Round struct {
	you, opponent string
}

func (r Round) outcomeScore() int {
	if r.you == same[r.opponent] {
		return 3
	}
	if r.you == beats[r.opponent] {
		return 6
	}
	return 0
}

func (r Round) score() int {
	return points[r.you] + r.outcomeScore()
}
