package main

import (
	"fmt"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	trees := readMap(lines)
	count := 0
	best := 0
	for x := 0; x < trees.width(); x++ {
		for y := 0; y < trees.height(); y++ {
			p := Point{x, y}
			if trees.visible(p) {
				count++
			}
			score := trees.score(p)
			if score > best {
				best = score
			}
		}
	}
	fmt.Println("visible from outside", count)
	fmt.Println("best score", best)
}

type treeMap [][]int

func (t treeMap) width() int {
	return len(t[0])
}

func (t treeMap) height() int {
	return len(t)
}

func (t treeMap) at(p Point) int {
	return t[p.y][p.x]
}

func (t treeMap) contains(p Point) bool {
	return p.x >= 0 && p.x < t.width() && p.y >= 0 && p.y < t.height()
}

func (t treeMap) visibleInDir(p, dir Point) (bool, int) {
	val := t.at(p)
	cur := p.add(dir)
	seen := 1
	for t.contains(cur) {
		if t.at(cur) >= val {
			return false, seen
		}
		cur = cur.add(dir)
		seen++
	}
	return true, seen - 1 // -1 because we fell off the map.
}

func (t treeMap) visible(p Point) bool {
	for _, dir := range directions {
		v, _ := t.visibleInDir(p, dir)
		if v {
			return true
		}
	}
	return false
}

func (t treeMap) score(p Point) int {
	result := 1
	for _, dir := range directions {
		_, seen := t.visibleInDir(p, dir)
		result *= seen
	}
	return result
}

var directions = []Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func readMap(lines []string) treeMap {
	result := make(treeMap, len(lines))
	for i, line := range lines {
		result[i] = make([]int, len(line))
		for j, c := range line {
			result[i][j] = lib.ParseInt(string(c))
		}
	}
	return result
}

type Point struct {
	x, y int
}

func (p Point) add(o Point) Point {
	return Point{p.x + o.x, p.y + o.y}
}
