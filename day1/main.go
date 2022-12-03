package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	file, err := os.Open("day1.txt")
	lib.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	var vals []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && total != 0 {
			vals = append(vals, total)
			total = 0
			continue
		}
		val, err := strconv.Atoi(line)
		lib.Check(err)
		total += val
	}

	sort.Ints(vals)
	fmt.Println("highest", vals[len(vals)-1])
	sum := 0
	for _, v := range vals[len(vals)-3:] {
		sum += v
	}
	fmt.Println("top 3 sum", sum)
}
