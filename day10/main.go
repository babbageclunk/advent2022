package main

import (
	"fmt"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	sampleTimes := lib.NewSet(20, 60, 100, 140, 180, 220)
	cpu := CPU{
		X:           1,
		SampleTimes: sampleTimes,
	}
	lines := lib.ReadLines("input")
	for _, line := range lines {
		cpu.apply(line)
	}
	fmt.Println(cpu.sum())
}

type CPU struct {
	X    int
	Time int

	SampleTimes  lib.Set[int]
	SampleValues []int
	Pixels       []rune
}

func (c *CPU) advance(n int) {
	for i := 0; i < n; i++ {
		c.Time++
		c.pixel()
		if c.SampleTimes.Has(c.Time) {
			c.sample()
		}
	}
}

func (c *CPU) sample() {
	c.SampleValues = append(c.SampleValues, c.X*c.Time)
}

func (c *CPU) pixel() {
	currentX := len(c.Pixels)
	current := '.'
	if c.X-1 <= currentX && currentX <= c.X+1 {
		current = '#'
	}
	c.Pixels = append(c.Pixels, current)
	if len(c.Pixels) == 40 {
		fmt.Println(string(c.Pixels))
		c.Pixels = nil
	}
}

func (c *CPU) sum() int {
	total := 0
	for _, v := range c.SampleValues {
		total += v
	}
	return total
}

func (c *CPU) apply(line string) {
	parts := strings.Split(line, " ")
	switch parts[0] {
	case "noop":
		c.advance(1)
	case "addx":
		i := lib.ParseInt(parts[1])
		c.advance(2)
		c.X += i
	}
}
