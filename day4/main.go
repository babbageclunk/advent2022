package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		a := ParseRange(parts[0])
		b := ParseRange(parts[1])
		if a.ContainsRange(b) || b.ContainsRange(a) {
			total++
		}
	}
	fmt.Println(total)
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	return Range{
		start: lib.ParseInt(parts[0]),
		end:   lib.ParseInt(parts[1]),
	}
}

type Range struct {
	start, end int
}

func (r Range) Contains(n int) bool {
	return r.start <= n && n <= r.end
}

func (r Range) ContainsRange(other Range) bool {
	return r.Contains(other.start) && r.Contains(other.end)
}
