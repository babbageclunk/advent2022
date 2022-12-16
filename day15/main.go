package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/samber/lo"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	// w := NewWorld(lib.ReadLines("sample-input"))
	w := NewWorld(lib.ReadLines("input"))
	count := 0
	// var line []rune
	for x := w.minx; x <= w.maxx; x++ {
		val := w.check(lib.Pt(x, 2000000))
		// c := '.'
		if val {
			count++
			// c = '#'
		}
		// line = append(line, c)
	}
	fmt.Println(count)
	// fmt.Println(string(line))
}

func NewWorld(lines []string) *World {
	world := World{
		sensors: make(map[lib.Point]int),
		beacons: lib.NewSet[lib.Point](),
		minx:    math.MaxInt,
		miny:    math.MaxInt,
		maxx:    math.MinInt,
		maxy:    math.MinInt,
	}
	lo.ForEach(lines, func(line string, _ int) {
		world.addBeacon(line)
	})
	return &world
}

type World struct {
	sensors                map[lib.Point]int
	beacons                lib.Set[lib.Point]
	minx, miny, maxx, maxy int
}

func (w *World) updateBounds(pt lib.Point) {
	if pt.X < w.minx {
		w.minx = pt.X
	}
	if pt.Y < w.miny {
		w.miny = pt.Y
	}
	if pt.X > w.maxx {
		w.maxx = pt.X
	}
	if pt.Y > w.maxy {
		w.maxy = pt.Y
	}
}

func coord(item string) int {
	return lib.ParseInt(strings.TrimRight(item[2:], ",:"))
}

func (w *World) addBeacon(line string) {
	parts := strings.Split(line, " ")
	sensor := lib.Pt(coord(parts[2]), coord(parts[3]))
	beacon := lib.Pt(coord(parts[8]), coord(parts[9]))
	dist := sensor.Sub(beacon).Manhattan()
	w.sensors[sensor] = dist
	w.beacons.Add(beacon)
	w.updateBounds(sensor.Add(lib.Pt(dist, dist)))
	w.updateBounds(sensor.Sub(lib.Pt(dist, dist)))
	w.updateBounds(beacon)
}

func (w *World) check(pt lib.Point) bool {
	if w.beacons.Has(pt) {
		return false
	}
	for sensor, dist := range w.sensors {
		if sensor.Sub(pt).Manhattan() <= dist {
			return true
		}
	}
	return false
}
