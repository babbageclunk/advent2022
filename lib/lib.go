package lib

import (
	"io/ioutil"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
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

func NewSet[T comparable](vals ...T) Set[T] {
	result := make(Set[T])
	for _, v := range vals {
		result.Add(v)
	}
	return result
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

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

type Point struct {
	X, Y int
}

func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

func (p Point) Sub(o Point) Point {
	return Pt(p.X-o.X, p.Y-o.Y)
}

func (p Point) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}

func ParsePoint(s string) Point {
	parts := strings.Split(s, ",")
	return Point{ParseInt(parts[0]), ParseInt(parts[1])}
}

var Directions = []Point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Sign(val int) int {
	switch {
	case val < 0:
		return -1
	case val > 0:
		return 1
	default:
		return 0
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Permuter[T any] struct {
	items []T
	perm  []int
}

func NewPermuter[T any](items []T) *Permuter[T] {
	return &Permuter[T]{items: items, perm: make([]int, len(items))}
}

func (p *Permuter[T]) More() bool {
	return p.perm[0] < len(p.perm)
}

func (p *Permuter[T]) Next() {
	for i := len(p.perm) - 1; i >= 0; i-- {
		if i == 0 || p.perm[i] < len(p.perm)-i-1 {
			p.perm[i]++
			return
		}
		p.perm[i] = 0
	}
}

func (p *Permuter[T]) Current() []T {
	result := make([]T, len(p.items))
	copy(result, p.items)
	for i, v := range p.perm {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func AllPermutations[T any](items []T) [][]T {
	var result [][]T
	for p := NewPermuter(items); p.More(); p.Next() {
		result = append(result, p.Current())
	}
	return result
}

type Queue[T any] struct {
	len        int
	head, tail *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func NewQueue[T any](items ...T) *Queue[T] {
	q := &Queue[T]{len: 0, head: nil, tail: nil}
	for _, item := range items {
		q.Push(item)
	}
	return q
}

func (q *Queue[T]) Len() int {
	return q.len
}

func (q *Queue[T]) Push(v T) {
	n := &node[T]{value: v}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		n.prev = q.tail
		q.tail = n
	}
	q.len++
}

func (q *Queue[T]) Pop() T {
	if q.head == nil {
		panic("empty queue")
	}
	result := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	} else {
		q.head.prev = nil
	}
	q.len--
	return result
}

func (q *Queue[T]) Peek() T {
	if q.head == nil {
		panic("empty queue")
	}
	return q.head.value
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.len = 0
}

func (q *Queue[T]) ToSlice() []T {
	result := make([]T, q.len)
	for i, n := 0, q.head; n != nil; i, n = i+1, n.next {
		result[i] = n.value
	}
	return result
}

func CopyAppend[T any](items []T, item T) []T {
	result := make([]T, len(items)+1)
	copy(result, items)
	result[len(items)] = item
	return result
}
