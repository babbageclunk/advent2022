package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	jets := []rune(strings.TrimSpace(lib.Read("input")))
	nudges := lo.Map(jets, func(j rune, _ int) lib.Point {
		if j == '<' {
			return left
		}
		return right
	})
	nudger := Nudger{nudges: nudges, pos: 0}
	var t Tower
	for i := 0; i < 2022; i++ {
		sprite := sprites[i%len(sprites)]
		pos := lib.Pt(2, t.highestBlock()+3)
		// fmt.Println("dropping", sprite, "at", pos)
		t.grow((pos.Y + sprite.height()) - t.top())
		for pos.Y >= 0 {
			newPos := pos
			dir := nudger.next()
			nudged := newPos.Add(dir)
			inBounds := nudged.X >= 0 && nudged.X+sprite.width() <= towerWidth
			if inBounds && sprite.canDraw(&t, nudged.X, nudged.Y) {
				// fmt.Println("nudged by", dir)
				newPos = nudged
			}
			dropped := newPos.Add(down)
			if dropped.Y < 0 || !sprite.canDraw(&t, dropped.X, dropped.Y) {
				// Keep any successful nudge
				pos = newPos
				break
			}
			// fmt.Println("dropping")
			pos = dropped
		}
		sprite.draw(&t, pos.X, pos.Y)
	}
	fmt.Println(t.highestBlock())
	// t.dump()
}

var (
	left  = lib.Directions[0]
	down  = lib.Directions[1]
	right = lib.Directions[2]
)

const (
	towerWidth   = 7
	maxBlocks    = 2022
	leftOffset   = 2
	bottomOffset = 3
)

type Nudger struct {
	nudges []lib.Point
	pos    int
}

func (n *Nudger) next() lib.Point {
	res := n.nudges[n.pos]
	n.pos = (n.pos + 1) % len(n.nudges)
	return res
}

type Tower struct {
	lines [][]rune
}

func (t *Tower) dump() {
	for _, row := range reversed(t.lines) {
		fmt.Printf("#%s#\n", string(row))
	}
	fmt.Println("#########")
}

func (t *Tower) grow(n int) {
	for i := 0; i < n; i++ {
		t.lines = append(t.lines, []rune("       "))
	}
}

func (t *Tower) top() int {
	return len(t.lines)
}

func (t *Tower) highestBlock() int {
	for i := len(t.lines) - 1; i >= 0; i-- {
		for _, c := range t.lines[i] {
			if c != ' ' {
				return i + 1
			}
		}
	}
	return 0
}

var sprites = makeSprites([][]string{{
	"1111",
}, {
	" 2 ",
	"222",
	" 2 ",
}, {
	"  3",
	"  3",
	"333",
}, {
	"4",
	"4",
	"4",
	"4",
}, {
	"55",
	"55",
}})

func makeSprites(items [][]string) []Sprite {
	result := make([]Sprite, len(items))
	for i, item := range items {
		result[i] = make(Sprite, len(item))
		for j, line := range lo.Reverse(item) {
			result[i][j] = []rune(line)
		}
	}
	return result
}

type Sprite [][]rune

func (s Sprite) width() int {
	return len(s[0])
}

func (s Sprite) height() int {
	return len(s)
}

func (s Sprite) draw(t *Tower, x, y int) {
	for i, row := range s {
		for j, c := range row {
			if c == ' ' {
				continue
			}
			t.lines[y+i][x+j] = c
		}
	}
}

func (s Sprite) canDraw(t *Tower, x, y int) bool {
	for i, row := range s {
		for j, c := range row {
			if c == ' ' {
				continue
			}
			if t.lines[y+i][x+j] != ' ' {
				return false
			}
		}
	}
	return true
}

func reversed(lines [][]rune) [][]rune {
	result := make([][]rune, len(lines))
	for i, line := range lines {
		result[len(lines)-i-1] = line
	}
	return result
}
