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
	return strings.Split(strings.TrimSpace(data), "\n")
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
