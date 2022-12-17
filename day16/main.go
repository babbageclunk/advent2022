package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
	"github.com/samber/lo"
)

func main() {
	g := make(Graph)
	for _, line := range lib.ReadLines("sample-input") {
		node := readNode(line)
		g[node.name] = node
	}
	g.dumpDot()
	// dist := g.shortestDist("AA")
	// fmt.Println(dist)
	// fmt.Println(g.route("AA", "JJ", dist))
	fmt.Println(findMaxFlow(g))
}

func findMaxFlow(g Graph) int {
	s := Solver{
		pos:    "AA",
		graph:  g,
		valves: make(map[string]int),
		path:   nil,
		time:   0,
	}
	total := 0
	for t := 1; t <= totalTime; t++ {
		s.step()
		s.time = t
		total += s.flow()
	}
	return total
}

const totalTime = 30

type Solver struct {
	pos    string
	graph  Graph
	valves map[string]int
	path   []string
	time   int
}

func (s *Solver) step() {
	// If we're walking, take the next move
	if len(s.path) > 0 {
		fmt.Printf("Walking from %s to %s (towards %s)\n", s.pos, s.path[0], s.path[len(s.path)-1])
		s.pos = s.path[0]
		s.path = s.path[1:]
		return
	}
	// Work out the best valve not-currently-on to turn on.
	dist := s.graph.shortestDist(s.pos)
	notOn := lo.Filter(lo.Values(s.graph), func(n *Node, _ int) bool {
		return s.valves[n.name] == 0
	})
	best := lo.MaxBy(notOn, func(a, b *Node) bool {
		remaining := totalTime - s.time
		aval := a.rate * (remaining - dist[a.name])
		bval := b.rate * (remaining - dist[b.name])
		return aval > bval
	})
	if best.name == s.pos {
		fmt.Printf("Turning on %s\n", best.name)
		s.valves[best.name] = best.rate
		return
	}
	s.path = s.graph.route(s.pos, best.name, dist)
	fmt.Printf("Next valve is %s, path %s\n", best.name, s.path)
}

func (s *Solver) flow() int {
	return lo.Sum(lo.Values(s.valves))
}

type Node struct {
	name       string
	rate       int
	neighbours []string
}

type Graph map[string]*Node

func readNode(line string) *Node {
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	parts := strings.Split(line, " ")
	rate := lib.ParseInt(strings.TrimPrefix(strings.TrimRight(parts[4], ";"), "rate="))
	return &Node{
		name: parts[1],
		rate: rate,
		neighbours: lo.Map(parts[9:], func(s string, _ int) string {
			return strings.TrimRight(s, ",")
		}),
	}
}

func (g Graph) dumpDot() {
	outf, err := os.Create("graph.dot")
	lib.Check(err)
	defer outf.Close()
	fmt.Fprintln(outf, "digraph {")
	for _, node := range g {
		fmt.Fprintf(outf, "%s [label=\"%s\\n%d\"];\n", node.name, node.name, node.rate)
		for _, neighbour := range node.neighbours {
			fmt.Fprintf(outf, "%s -> %s;\n", node.name, neighbour)
		}
	}
	fmt.Fprintln(outf, "}")
}

func (g Graph) shortestDist(start string) map[string]int {
	// Use Dijkstra's algorithm to find the shortest path from start
	// to every other node.
	unvisited := lib.NewSet(lo.Keys(g)...)
	dist := lo.MapEntries(g, func(key string, _ *Node) (string, int) {
		return key, math.MaxInt
	})
	dist[start] = 0
	current := start

	for {
		neigbours := lo.Filter(g[current].neighbours, func(k string, _ int) bool {
			return unvisited.Has(k)
		})
		for _, neighbour := range neigbours {
			dist[neighbour] = lib.Min(dist[neighbour], dist[current]+1)
		}
		unvisited.Remove(current)
		if unvisited.Len() == 0 {
			return dist
		}
		current = lo.MinBy(lo.Keys(unvisited), func(a, b string) bool {
			return dist[a] < dist[b]
		})
	}
}

func (g Graph) route(start, end string, dist map[string]int) []string {
	// Assuming all links are bidirectional, find the shortest route
	// given the distances.
	route := []string{end}
	current := end
	for {
		neighbours := lo.Filter(g[current].neighbours, func(k string, _ int) bool {
			return dist[k] == dist[current]-1
		})
		current = neighbours[0]
		if current == start {
			break
		}
		route = append(route, current)
	}
	return lo.Reverse(route)
}
