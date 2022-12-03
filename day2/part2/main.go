package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data, err := ioutil.ReadFile("../input")
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
			outcome:  parts[1],
		}.score()
	}
	fmt.Println(total)
}

// Rock - A X
// Paper - B Y
// Scissors - C Z
var beats map[string]string = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

var points map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var outcomeScore map[string]int = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

type Round struct {
	opponent, outcome string
}

func (r Round) you() string {
	switch r.outcome {
	case "X": // Lose
		return beats[beats[r.opponent]]
	case "Y": // Draw
		return r.opponent
	case "Z": // Win
		return beats[r.opponent]
	}
	panic("uh oh")
}

func (r Round) score() int {
	return points[r.you()] + outcomeScore[r.outcome]
}
