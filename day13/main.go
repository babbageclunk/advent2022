package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	if len(os.Args) < 2 || os.Args[1] == "part1" {
		part1(lines)
	} else {
		part2(lines)
	}
}

func part1(lines []string) {
	ix := 1
	total := 0
	for i := 0; i < len(lines); i += 3 {
		left := read(lines[i])
		right := read(lines[i+1])
		res := cmpListList(left, right)
		// fmt.Println(ix, left, right, res)
		if res == 1 {
			total += ix
		}
		ix++
	}
	fmt.Println(total)
}

func part2(lines []string) {
	m1 := read("[[2]]")
	m2 := read("[[6]]")
	packets := []any{m1, m2}
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets = append(packets, read(line))
	}
	sort.Slice(packets, func(i, j int) bool {
		return cmp(packets[i], packets[j]) >= 0
	})
	var m1Pos, m2Pos int
	for i, p := range packets {
		// fmt.Println(p)
		if cmp(m1, p) == 0 {
			m1Pos = i + 1
		}
		if cmp(m2, p) == 0 {
			m2Pos = i + 1
		}
	}
	fmt.Println(m1Pos * m2Pos)
}

func read(value string) []any {
	var items []any
	err := json.Unmarshal([]byte(value), &items)
	lib.Check(err)
	return items
}

func cmpListList(left, right []any) int {
	for i, val := range left {
		if i >= len(right) {
			return -1
		}
		res := cmp(val, right[i])
		if res != 0 {
			return res
		}
	}
	if len(left) < len(right) {
		return 1
	}
	return 0
}

func cmp(left, right any) int {
	switch v := left.(type) {
	case float64:
		return cmpInt(v, right)
	case []any:
		return cmpList(v, right)
	default:
		panic(fmt.Sprintf("unknown type %T", left))
	}
}

func cmpInt(left float64, right any) int {
	switch v := right.(type) {
	case float64:
		return sign(v - left)
	case []any:
		return cmpListList([]any{left}, v)
	default:
		panic(fmt.Sprintf("unknown type %T", right))
	}
}

func cmpList(left []any, right any) int {
	switch v := right.(type) {
	case float64:
		return cmpListList(left, []any{v})
	case []any:
		return cmpListList(left, v)
	default:
		panic(fmt.Sprintf("unknown type %T", right))
	}
}

func sign(val float64) int {
	switch {
	case val > 0:
		return 1
	case val < 0:
		return -1
	default:
		return 0
	}
}
