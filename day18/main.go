package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
	"github.com/samber/lo"
)

func main() {
	input := "sample-input"
	if len(os.Args) > 2 {
		input = os.Args[2]
	}
	lines := lib.ReadLines(input)
	if len(os.Args) < 2 || os.Args[1] == "part1" {
		part1(lines)
	} else {
		part2(lines)
	}
}

var directions = []Point3{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

type Point3 struct {
	x, y, z int
}

func (p Point3) add(other Point3) Point3 {
	return Point3{p.x + other.x, p.y + other.y, p.z + other.z}
}

func (p Point3) neighbours() []Point3 {
	return lo.Map(directions, func(d Point3, _ int) Point3 {
		return p.add(d)
	})
}

func readPoint(line string) Point3 {
	parts := strings.Split(line, ",")
	return Point3{
		x: lib.ParseInt(parts[0]),
		y: lib.ParseInt(parts[1]),
		z: lib.ParseInt(parts[2]),
	}
}

type Graph map[Point3]lib.Set[Point3]

func (g Graph) add(p Point3) {
	neighbours := lib.NewSet[Point3]()
	for _, neighbour := range p.neighbours() {
		otherNeighbours, found := g[neighbour]
		if !found {
			continue
		}
		neighbours.Add(neighbour)
		otherNeighbours.Add(p)
	}
	g[p] = neighbours
}

func (g Graph) neighbours(p Point3) []Point3 {
	var result []Point3
	for _, n := range p.neighbours() {
		if _, found := g[n]; found {
			result = append(result, n)
		}
	}
	return result
}

func readGraph(lines []string) Graph {
	graph := make(Graph, len(lines))
	for _, line := range lines {
		pt := readPoint(line)
		graph.add(pt)
	}
	return graph
}

func part1(lines []string) {
	graph := readGraph(lines)
	// Surface area of the nodes is sum of (6 - edges) for each node.
	total := lo.SumBy(lo.Values(graph), func(item lib.Set[Point3]) int {
		return 6 - item.Len()
	})
	fmt.Println(total)
}

type GetterFunc func(Point3) int

var getters = []GetterFunc{
	func(p Point3) int { return p.x },
	func(p Point3) int { return p.y },
	func(p Point3) int { return p.z },
}

func allPoints(mins []int, maxs []int) []Point3 {
	total := (maxs[0] - mins[0]) * (maxs[1] - mins[1]) * (maxs[2] - mins[2])
	points := make([]Point3, 0, total)
	for x := mins[0]; x <= maxs[0]; x++ {
		for y := mins[1]; y <= maxs[1]; y++ {
			for z := mins[2]; z <= maxs[2]; z++ {
				points = append(points, Point3{x, y, z})
			}
		}
	}
	return points
}

func part2(lines []string) {
	graph := readGraph(lines)
	mins := lo.Map(getters, func(getter GetterFunc, _ int) int {
		return lo.Min(lo.Map(lo.Keys(graph), func(p Point3, _ int) int {
			return getter(p)
		}))
	})
	maxs := lo.Map(getters, func(getter GetterFunc, _ int) int {
		return lo.Max(lo.Map(lo.Keys(graph), func(p Point3, _ int) int {
			return getter(p)
		}))
	})

	fmt.Println(mins)
	fmt.Println(maxs)

	// Construct a parallel graph of the gaps between the cubes.
	gaps := make(Graph)
	for _, pt := range allPoints(mins, maxs) {
		if _, found := graph[pt]; found {
			continue
		}
		gaps.add(pt)
	}

	// Find all the gaps that are connected to the outside.
	connected := lib.NewSet[Point3]()
	unknown := lib.NewSet(lo.Keys(gaps)...)

	for unknown.Len() > 0 {
		cur := popOne(unknown)
		if connected.Has(cur) {
			continue
		}
		if outside(cur, mins, maxs) {
			connected.Add(cur)
			continue
		}
		todo := lib.NewSet(cur)
		done := lib.NewSet[Point3]()

		for todo.Len() > 0 {
			item := popOne(todo)
			for _, neighbour := range gaps.neighbours(item) {

				if connected.Has(neighbour) {
					// A neighbour is connected to the outside, so we are
					// too.
					connected.Add(item)
					continue
				}
				if outside(neighbour, mins, maxs) {
					connected.Add(neighbour)
					connected.Add(item)
					continue
				}
				if !done.Has(neighbour) {
					todo.Add(neighbour)
				}
			}
			done.Add(item)
		}
	}
	// connected should only contain the gaps that are connected to
	// the outside.
	fmt.Println("connected", connected.Len())
	// left, _ := lo.Difference(lo.Keys(gaps), lo.Keys(connected))
	// fmt.Println("left", left)

	total := 0
	for _, pt := range lo.Keys(graph) {
		for _, neighbour := range pt.neighbours() {
			if _, found := graph[neighbour]; found {
				continue
			}
			if outside(neighbour, mins, maxs) || connected.Has(neighbour) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func popOne[T comparable](s lib.Set[T]) T {
	for key := range s {
		s.Remove(key)
		return key
	}
	panic("empty set")
}

func outside(pt Point3, mins, maxs []int) bool {
	for i, getter := range getters {
		if getter(pt) <= mins[i] || getter(pt) >= maxs[i] {
			return true
		}
	}
	return false
}
