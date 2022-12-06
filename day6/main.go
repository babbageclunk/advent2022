package main

import (
	"fmt"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data := lib.Read("input")
	fmt.Println(detectStart(data, 4))
	fmt.Println(detectStart(data, 14))
}

func detectStart(data string, size int) int {
	runes := []rune(data)
	startChunk := runes[:size]
	buf := lib.NewRingBufferFrom(startChunk)
	distinct := lib.NewBagFrom(startChunk)
	pos := len(startChunk)
	for pos < len(runes) {
		if distinct.Len() == size {
			return pos
		}
		new := runes[pos]
		old := buf.Push(new)
		distinct.Add(new)
		distinct.Remove(old)
		pos++
	}
	panic("never found start")
}
