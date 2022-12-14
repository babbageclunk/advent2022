package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	world := World{
		spots: make(map[lib.Point]rune),
		minx:  math.MaxInt,
		miny:  math.MaxInt,
		maxx:  math.MinInt,
		maxy:  math.MinInt,
	}
	for _, line := range lines {
		world.draw(line)
	}
	grains := 0
	for world.drop() {
		grains++
	}
	fmt.Println(grains)
}

type World struct {
	spots                  map[lib.Point]rune
	minx, miny, maxx, maxy int
}

func (w *World) set(pt lib.Point, val rune) {
	w.spots[pt] = val
	if pt.X < w.minx {
		w.minx = pt.X
	}
	if pt.Y < w.miny {
		w.miny = pt.Y
	}
	if pt.X > w.maxx {
		w.maxx = pt.X
	}
	if pt.Y > w.maxy {
		w.maxy = pt.Y
	}
}

func (w *World) at(pt lib.Point) bool {
	_, found := w.spots[pt]
	return found
}

func points(x1, y1, x2, y2 int) []lib.Point {
	var pts []lib.Point
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			pts = append(pts, lib.Pt(x, y))
		}
	}
	return pts
}

func (w *World) draw(line string) {
	parts := strings.Split(line, " -> ")
	for i := 1; i < len(parts); i++ {
		p1 := lib.ParsePoint(parts[i-1])
		p2 := lib.ParsePoint(parts[i])
		for _, pt := range points(p1.X, p1.Y, p2.X, p2.Y) {
			w.set(pt, '#')
		}
	}
}

func (w *World) dump() {
	fmt.Println(w)
	for y := w.miny; y <= w.maxy; y++ {
		line := make([]rune, w.maxx-w.minx)
		for i, x := 0, w.minx; x < w.maxx; i, x = i+1, x+1 {
			val, found := w.spots[lib.Pt(x, y)]
			if !found {
				val = '.'
			}
			line[i] = val
		}
		fmt.Println(string(line))
	}
}

var (
	dirs = []lib.Point{
		{X: 0, Y: 1},  // down - Y is reversed.
		{X: -1, Y: 1}, // down-left
		{X: 1, Y: 1},  // down-right
	}
	source = lib.Pt(500, 0)
)

func (w *World) drop() bool {
	cur := source
	for cur.Y < w.maxy {
		moved := false
		for _, dir := range dirs {
			newPos := cur.Add(dir)
			if !w.at(newPos) {
				cur = newPos
				moved = true
				break
			}
		}
		if !moved {
			w.set(cur, 'o')
			return true
		}
	}
	// Fell off into the void!
	return false
}
