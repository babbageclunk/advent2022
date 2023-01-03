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

func part1(lines []string) {
	graph := make(map[Point3]lib.Set[Point3], len(lines))
	for _, line := range lines {
		pt := readPoint(line)
		neighbours := lib.NewSet[Point3]()
		for _, neighbour := range pt.neighbours() {
			otherNeighbours, found := graph[neighbour]
			if !found {
				continue
			}
			neighbours.Add(neighbour)
			otherNeighbours.Add(pt)
		}
		graph[pt] = neighbours
	}
	// Surface area of the nodes is sum of (6 - edges) for each node.
	total := lo.SumBy(lo.Values(graph), func(item lib.Set[Point3]) int {
		return 6 - item.Len()
	})
	fmt.Println(total)
}

func part2(lines []string) {}
