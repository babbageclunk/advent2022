package main

import (
	"fmt"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	data := lib.Read("input")
	fmt.Println(detectStart(data))
}

func detectStart(data string) int {
	runes := []rune(data)
	startChunk := runes[:4]
	buf := lib.NewRingBufferFrom(startChunk)
	distinct := lib.NewBagFrom(startChunk)
	pos := len(startChunk)
	for pos < len(runes) {
		if distinct.Len() == 4 {
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
