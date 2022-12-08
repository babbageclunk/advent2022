package main

import (
	"fmt"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	trees := readMap(lines)
	count := 0
	for x := 0; x < trees.width(); x++ {
		for y := 0; y < trees.height(); y++ {
			if trees.visible(Point{x, y}) {
				count++
			}
		}
	}
	fmt.Println(count)
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
	return p.x >= 0 && p.x < len(t[0]) && p.y >= 0 && p.y < len(t)
}

func (t treeMap) visibleInDir(p, dir Point) bool {
	val := t.at(p)
	cur := p.add(dir)
	for t.contains(cur) {
		if t.at(cur) >= val {
			return false
		}
		cur = cur.add(dir)
	}
	return true
}

func (t treeMap) visible(p Point) bool {
	for _, dir := range directions {
		if t.visibleInDir(p, dir) {
			return true
		}
	}
	return false
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
