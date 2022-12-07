package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	cursor := Cursor{lines: lines}
	// Skip the "cd /" line
	cursor.next()
	cur := makeNode("/")
	root := cur
loop:
	for {
		cmd := cursor.current()
		parts := strings.Split(cmd, " ")
		if parts[0] != "$" {
			panic(fmt.Sprintf("expected cmd at line %d, got %q", cursor.pos, cmd))
		}
		switch parts[1] {
		case "cd":
			if parts[2] == ".." {
				cur = cur.parent
			} else {
				cur = cur.dirs[parts[2]]
			}
			if !cursor.next() {
				break loop
			}
		case "ls":
			if !readDir(cur, &cursor) {
				break loop
			}
		}
	}

	total := 0
	for _, n := range root.allNodes() {
		size := n.size()
		if size <= 100000 {
			total += size
		}
	}
	fmt.Println(total)

	available := 70000000 - root.size()
	needed := 30000000 - available
	fmt.Println("need", needed)

	candidate := 70000000
	for _, n := range root.allNodes() {
		size := n.size()
		if size >= needed && size < candidate {
			candidate = size
		}
	}
	fmt.Println("should delete", candidate)
}

type Cursor struct {
	lines []string
	pos   int
}

func (c *Cursor) next() bool {
	if c.pos+1 == len(c.lines) {
		return false
	}
	c.pos++
	return true
}

func (c *Cursor) current() string {
	return c.lines[c.pos]
}

type Node struct {
	name   string
	parent *Node
	files  map[string]int
	dirs   map[string]*Node
}

func (n *Node) size() int {
	if n == nil {
		return 0
	}
	total := 0
	for _, size := range n.files {
		total += size
	}
	for _, child := range n.dirs {
		total += child.size()
	}
	return total
}

func (n *Node) allNodes() []*Node {
	result := []*Node{n}
	for _, child := range n.dirs {
		result = append(result, child.allNodes()...)
	}
	return result
}

func makeNode(name string) *Node {
	return &Node{
		name:   name,
		parent: nil,
		files:  make(map[string]int),
		dirs:   make(map[string]*Node),
	}
}

func readDir(node *Node, cursor *Cursor) bool {
	for cursor.next() {
		parts := strings.Split(cursor.current(), " ")
		switch parts[0] {
		case "$":
			return true
		case "dir":
			name := parts[1]
			newNode := makeNode(name)
			newNode.parent = node
			node.dirs[name] = newNode
		default:
			node.files[parts[1]] = lib.ParseInt(parts[0])
		}
	}
	return false
}
