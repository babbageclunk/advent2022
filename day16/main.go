package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/samber/lo"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	g := make(Graph)
	for _, line := range lib.ReadLines("input") {
		node := readNode(line)
		g[node.name] = node
	}
	g.dumpDot()
	if len(os.Args) < 2 || os.Args[1] == "part1" {
		fmt.Println(findMaxFlow(g, 30, &singleBest{}))
	} else {
		fmt.Println(findMaxFlow(g, 26, &bestTwo{
			best: make(map[string]Step),
		}))
	}
}

func findMaxFlow(g Graph, maxTime int, chooser StepChooser) int {
	dist := g.allDistances()
	nextSteps := lib.NewQueue(Step{
		node:      "AA",
		remaining: maxTime,
		total:     0,
		valves:    nil,
	})
	nonZero := lo.FilterMap(lo.Values(g), func(n *Node, _ int) (string, bool) {
		return n.name, n.rate != 0
	})
	rates := lo.MapEntries(g, func(k string, n *Node) (string, int) {
		return k, n.rate
	})
	steps := 0
	for nextSteps.Len() > 0 {
		steps++
		step := nextSteps.Pop()
		// Try every node that we haven't visited yet that we can get
		// to in the remaining time.
		unvisited, _ := lo.Difference(nonZero, step.valves)
		for _, node := range unvisited {
			timeToTurnOn := dist[np{step.node, node}] + 1
			if timeToTurnOn > step.remaining {
				continue
			}
			newRemaining := step.remaining - timeToTurnOn
			newStep := Step{
				node:      node,
				remaining: newRemaining,
				total:     step.total + (rates[node] * newRemaining),
				valves:    lib.CopyAppend(step.valves, node),
			}
			nextSteps.Push(newStep)
			chooser.Consider(newStep)
		}
	}
	fmt.Println(steps)
	bestSteps := chooser.Best()
	for _, step := range bestSteps {
		fmt.Printf("%+v\n", step)
	}
	return lo.SumBy(bestSteps, func(s Step) int {
		return s.total
	})
}

type StepChooser interface {
	Consider(Step)
	Best() []Step
}

type singleBest struct {
	best Step
}

func (s *singleBest) Consider(step Step) {
	if step.total > s.best.total {
		s.best = step
	}
}

func (s *singleBest) Best() []Step {
	return []Step{s.best}
}

type bestTwo struct {
	best map[string]Step
}

func (b *bestTwo) Consider(step Step) {
	key := toKey(step.valves)
	if existing := b.best[key]; step.total > existing.total {
		b.best[key] = step
	}
}

func (b *bestTwo) Best() []Step {
	total := 0
	bestPair := make([]Step, 2)
	fmt.Println(len(b.best), "possible best steps")
	expandedKeys := lo.MapEntries(b.best, func(k string, _ Step) (string, []string) {
		return k, fromKey(k)
	})
	for lkey, left := range b.best {
		leftKey := expandedKeys[lkey]
		for rkey, right := range b.best {
			rightKey := expandedKeys[rkey]
			if len(lo.Intersect(leftKey, rightKey)) > 0 {
				continue
			}
			if left.total+right.total > total {
				total = left.total + right.total
				bestPair[0] = left
				bestPair[1] = right
			}
		}
	}
	return bestPair
}

func toKey(v []string) string {
	vCopy := append([]string(nil), v...)
	sort.Strings(vCopy)
	return strings.Join(vCopy, ",")
}

func fromKey(key string) []string {
	return strings.Split(key, ",")
}

type Step struct {
	node      string
	remaining int
	total     int
	valves    []string
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

type np struct {
	a, b string
}

func (g Graph) allDistances() map[np]int {
	// Using Floyd-Warshall's algorithm, find the shortest distance
	// between all pairs of nodes.
	dist := make(map[np]int)
	lookup := func(a, b string) int {
		d, found := dist[np{a, b}]
		if !found {
			return math.MaxInt
		}
		return d
	}
	for _, node := range g {
		dist[np{node.name, node.name}] = 0
		for _, neighbour := range node.neighbours {
			dist[np{node.name, neighbour}] = 1
		}
	}
	addWithInf := func(a, b int) int {
		if a == math.MaxInt || b == math.MaxInt {
			return math.MaxInt
		}
		return a + b
	}
	for k := range g {
		for i := range g {
			for j := range g {
				iToK := lookup(i, k)
				kToJ := lookup(k, j)
				iToJ := addWithInf(iToK, kToJ)
				if lookup(i, j) > iToJ {
					dist[np{i, j}] = iToJ
				}
			}
		}
	}
	return dist
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
