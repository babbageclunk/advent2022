package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	m := readMap(lib.ReadLines("input"))
	fmt.Println(m.findShortest())
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

func (m heightMap) at(p lib.Point) rune {
	return m.heights[p.Y][p.X]
}

type NodeMap [][]*Node

func (m NodeMap) neighbours(n *Node) []*Node {
	var result []*Node
	for _, dir := range lib.Directions {
		candidate := n.loc.Add(dir)
		if m.contains(candidate) {
			result = append(result, m.at(candidate))
		}
	}
	return result
}

func (m NodeMap) width() int {
	return len(m[0])
}

func (m NodeMap) height() int {
	return len(m)
}

func (m NodeMap) at(p lib.Point) *Node {
	return m[p.Y][p.X]
}

func (m NodeMap) contains(p lib.Point) bool {
	return p.X >= 0 && p.X < m.width() && p.Y >= 0 && p.Y < m.height()
}

type Node struct {
	loc  lib.Point
	dist int
}

func (m heightMap) findShortest() int {
	// Bastardised version of Dijkstra's algorithm.
	// Make nodemap
	unvisited := lib.NewSet[lib.Point]()
	nodes := make(NodeMap, len(m.heights))
	for y, row := range m.heights {
		nodes[y] = make([]*Node, len(row))
		for x := range row {
			node := Node{
				loc:  lib.Point{X: x, Y: y},
				dist: math.MaxInt,
			}
			nodes[y][x] = &node
			unvisited.Add(node.loc)
		}
	}
	start := nodes.at(m.start)
	start.dist = 0
	current := start

	dest := nodes.at(m.end)

	for {
		neighbours := nodes.neighbours(current)
		for _, n := range neighbours {
			if !unvisited.Has(n.loc) {
				continue
			}
			// This is only a neighbour if it's a valid step in the height map.
			if !validStep(m.at(current.loc), m.at(n.loc)) {
				continue
			}
			dist := current.dist + 1
			if dist < n.dist {
				n.dist = dist
			}
		}
		unvisited.Remove(current.loc)

		if !unvisited.Has(dest.loc) {
			break
		}
		var smallestUnvisited *Node
		for loc := range unvisited {
			node := nodes.at(loc)
			if smallestUnvisited == nil || node.dist < smallestUnvisited.dist {
				smallestUnvisited = node
			}
		}
		if smallestUnvisited == dest {
			break
		}
		current = smallestUnvisited
	}

	return dest.dist
}

func validStep(from, to rune) bool {
	return from+1 >= to
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
