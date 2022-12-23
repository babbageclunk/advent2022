package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	var t Tower
	t.grow(10)
	sprites[2].draw(&t, 0, 0)
	sprites[1].draw(&t, 0, 2)
	t.dump()

	fmt.Println(sprites[0].canDraw(&t, 0, 0))
	fmt.Println(sprites[0].canDraw(&t, 3, 0))
}

const (
	towerWidth   = 7
	maxBlocks    = 2022
	leftOffset   = 2
	bottomOffset = 3
)

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
