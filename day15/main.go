package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/samber/lo"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "part1" {
		part1()
	} else {
		part2()
	}
}

func part1() {
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

const scale = 4_000_000

func part2() {
	w := NewWorld(lib.ReadLines("input"))
	bset := lib.NewSet[lib.Point]()
	for pt, dist := range w.sensors {
		for _, b := range boundaries(pt, dist) {
			if b.X < 0 || b.X > scale {
				continue
			}
			if b.Y < 0 || b.Y > scale {
				continue
			}
			bset.Add(b)
		}
	}
	fmt.Println(len(bset))
pts:
	for pt := range bset {
		// fmt.Println(pt)
		// Check whether this point is further than dist from each sensor.
		for beacon, dist := range w.sensors {
			if beacon.Sub(pt).Manhattan() <= dist {
				continue pts
			}
		}
		// This must be far enough from all beacons.
		fmt.Println(pt, (4_000_000*pt.X)+pt.Y)
		break
	}
}

func boundaries(pt lib.Point, dist int) []lib.Point {
	var res []lib.Point
	// Add top and bottom
	d := lib.Pt(0, dist+1)
	res = append(res, pt.Sub(d), pt.Add(d))
	for delta := -dist; delta <= dist; delta++ {
		left := dist - lib.Abs(delta) + 1
		d1 := lib.Pt(pt.X-left, pt.Y+delta)
		d2 := lib.Pt(pt.X+left, pt.Y+delta)
		res = append(res, d1, d2)
	}
	return res
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
