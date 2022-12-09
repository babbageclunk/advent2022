package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	rope2 := NewRope(2)
	rope10 := NewRope(10)
	for _, line := range lib.ReadLines("input") {
		move := readMove(line)
		rope2.apply(move)
		rope10.apply(move)
	}
	fmt.Println(rope2.seen.Len())
	fmt.Println(rope10.seen.Len())
}

var dirNames = map[string]lib.Point{
	"L": lib.Directions[0],
	"D": lib.Directions[1],
	"R": lib.Directions[2],
	"U": lib.Directions[3],
}

func NewRope(n int) *Rope {
	return &Rope{
		parts: make([]lib.Point, n),
		seen:  lib.NewSet[lib.Point](),
	}
}

type Rope struct {
	parts []lib.Point
	seen  lib.Set[lib.Point]
}

func (r *Rope) apply(m Move) {
	for i := 0; i < m.steps; i++ {
		r.step(m.dir)
	}
}

func (r *Rope) step(dir lib.Point) {
	r.parts[0] = r.parts[0].Add(dir)
	r.moveTails()
}

func (r *Rope) moveTails() {
	for i := 1; i < len(r.parts); i++ {
		r.moveTail(i)
	}
	r.seen.Add(r.parts[len(r.parts)-1])
}

func (r *Rope) moveTail(n int) {
	prev := r.parts[n-1]
	this := r.parts[n]
	dx := prev.X - this.X
	dy := prev.Y - this.Y
	if abs(dx) < 2 && abs(dy) < 2 {
		return
	}
	move := lib.Point{
		X: sign(dx),
		Y: sign(dy),
	}
	r.parts[n] = this.Add(move)
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
