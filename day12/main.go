package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	m := readMap(lib.ReadLines("sample-input"))
	path := m.solveFrom(Path{m.start})
	fmt.Println(len(path) - 1)
	fmt.Println(path)
}

func readMap(lines []string) heightMap {
	result := heightMap{
		heights: make([][]rune, len(lines)),
	}
	for y, line := range lines {
		result.heights[y] = []rune(line)
		if x := strings.Index(line, "S"); x != -1 {
			result.heights[y][x] = 'a'
			result.start = lib.Point{X: x, Y: y}
		}
		if x := strings.Index(line, "E"); x != -1 {
			result.heights[y][x] = 'z'
			result.end = lib.Point{X: x, Y: y}
		}
	}
	return result
}

type heightMap struct {
	heights    [][]rune
	start, end lib.Point
}

func (m heightMap) width() int {
	return len(m.heights[0])
}

func (m heightMap) height() int {
	return len(m.heights)
}

func (m heightMap) at(p lib.Point) rune {
	return m.heights[p.Y][p.X]
}

func (m heightMap) contains(p lib.Point) bool {
	return p.X >= 0 && p.X < m.width() && p.Y >= 0 && p.Y < m.height()
}

func (m heightMap) neighbours(p lib.Point) []lib.Point {
	var result []lib.Point
	for _, dir := range lib.Directions {
		candidate := p.Add(dir)
		if m.contains(candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

func (m heightMap) solveFrom(path Path) Path {
	last := path[len(path)-1]
	height := m.at(last)
	if last == m.end {
		return path
	}
	var bestPath Path
	for _, n := range m.neighbours(last) {
		newHeight := m.at(n)
		if !validStep(height, newHeight) {
			continue
		}
		if newPath := path.add(n); newPath != nil {
			newRes := m.solveFrom(newPath)
			if newRes == nil {
				continue
			}
			if bestPath == nil || len(newRes) < len(bestPath) {
				bestPath = newRes
			}
		}
	}
	return bestPath
}

func validStep(from, to rune) bool {
	diff := to - from
	return diff == 0 || diff == 1
}

type Path []lib.Point

func (p Path) contains(c lib.Point) bool {
	for _, point := range p {
		if point == c {
			return true
		}
	}
	return false
}

func (p Path) add(c lib.Point) Path {
	if p.contains(c) {
		return nil
	}
	result := make(Path, len(p)+1)
	copy(result, p)
	result[len(p)] = c
	return result
}
