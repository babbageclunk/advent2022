package lib

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	Check(err)
	return string(data)
}

func ReadLines(filename string) []string {
	data := Read(filename)
	return strings.Split(strings.TrimRight(data, "\n"), "\n")
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	Check(err)
	return n
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) ToSlice() []T {
	var result []T
	for k := range s {
		result = append(result, k)
	}
	return result
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := NewSet[T]()
	for i := range s {
		if other.Has(i) {
			result.Add(i)
		}
	}
	return result
}

type Bag[T comparable] map[T]int

func NewBagFrom[T comparable](vs []T) Bag[T] {
	result := make(Bag[T])
	for _, v := range vs {
		result.Add(v)
	}
	return result
}

func (b Bag[T]) Add(v T) {
	b[v]++
}

func (b Bag[T]) Remove(v T) {
	if b[v] > 0 {
		b[v]--
	}
	if b[v] == 0 {
		delete(b, v)
	}
}

func (b Bag[T]) Len() int {
	return len(b)
}

type RingBuffer[T any] struct {
	items []T
	pos   int
}

func NewRingBufferFrom[T any](vs []T) *RingBuffer[T] {
	return &RingBuffer[T]{items: vs, pos: 0}
}

func (r *RingBuffer[T]) Push(v T) T {
	result := r.items[r.pos]
	r.items[r.pos] = v
	r.pos = (r.pos + 1) % len(r.items)
	return result
}

type Point struct {
	X, Y int
}

func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

var Directions = []Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}
