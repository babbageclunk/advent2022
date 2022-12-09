package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	rope := Rope{seen: lib.NewSet[lib.Point]()}
	rope.seen.Add(lib.Point{X: 0, Y: 0})
	for _, line := range lib.ReadLines("input") {
		move := readMove(line)
		rope.apply(move)
	}
	fmt.Println(rope.seen.Len())
}

var dirNames = map[string]lib.Point{
	"L": lib.Directions[0],
	"D": lib.Directions[1],
	"R": lib.Directions[2],
	"U": lib.Directions[3],
}

type Rope struct {
	head, tail lib.Point
	seen       lib.Set[lib.Point]
}

func (r *Rope) apply(m Move) {
	for i := 0; i < m.steps; i++ {
		r.step(m.dir)
	}
}

func (r *Rope) step(dir lib.Point) {
	r.head = r.head.Add(dir)
	r.moveTail()
}

func (r *Rope) moveTail() {
	dx := r.head.X - r.tail.X
	dy := r.head.Y - r.tail.Y
	if abs(dx) < 2 && abs(dy) < 2 {
		return
	}
	tailMove := lib.Point{
		X: sign(dx),
		Y: sign(dy),
	}
	r.tail = r.tail.Add(tailMove)
	r.seen.Add(r.tail)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func sign(val int) int {
	switch {
	case val < 0:
		return -1
	case val > 0:
		return 1
	default:
		return 0
	}
}

type Move struct {
	dir   lib.Point
	steps int
}

func readMove(line string) Move {
	parts := strings.Split(line, " ")
	return Move{
		dir:   dirNames[parts[0]],
		steps: lib.ParseInt(parts[1]),
	}
}
