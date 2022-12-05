package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	contained := 0
	overlaps := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		a := ParseRange(parts[0])
		b := ParseRange(parts[1])
		if a.ContainsRange(b) || b.ContainsRange(a) {
			contained++
		}
		if a.Overlaps(b) {
			overlaps++
		}
	}
	fmt.Println("contained", contained)
	fmt.Println("overlaps", overlaps)
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

func (r Range) Overlaps(other Range) bool {
	return r.Contains(other.start) || r.Contains(other.end) || r.ContainsRange(other) || other.ContainsRange(r)
}
