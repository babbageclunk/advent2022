package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/babbageclunk/advent2022/lib"
)

func main() {
	lines := lib.ReadLines("input")
	var monkeys []*Monkey
	for i := 0; i < len(lines); i += 7 {
		m := readMonkey(lines[i : i+6])
		monkeys = append(monkeys, m)
	}
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.run(monkeys)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	fmt.Println(monkeys[0].inspected * monkeys[1].inspected)
}

type Monkey struct {
	id        int
	items     []int
	operation Op
	next      Next
	inspected int
}

func (m *Monkey) run(monkeys []*Monkey) {
	for _, val := range m.items {
		newVal := m.operation.apply(val)
		newVal /= 3
		next := m.next.apply(newVal)
		monkeys[next].items = append(monkeys[next].items, newVal)
		m.inspected++
	}
	m.items = nil
}

func readMonkey(lines []string) *Monkey {
	if len(lines) != 6 {
		panic(fmt.Sprintf("monkey has wrong # of lines: %q", lines))
	}
	return &Monkey{
		id:        readId(lines[0]),
		items:     readItems(lines[1]),
		operation: readOperation(lines[2]),
		next:      readNext(lines[3:]),
	}
}

var idPattern = regexp.MustCompile(`Monkey (\d+):`)

func readId(line string) int {
	matches := idPattern.FindStringSubmatch(line)
	if len(matches) != 2 {
		panic(fmt.Sprintf("%q doesn't match id pattern", line))
	}
	return lib.ParseInt(matches[1])
}

func checkPrefix(prefix, line string) string {
	if !strings.HasPrefix(line, prefix) {
		panic(fmt.Sprintf("%q doesn't match prefix %q", line, prefix))
	}
	return line[len(prefix):]
}

const itemsPrefix = "  Starting items: "

func readItems(line string) []int {
	parts := strings.Split(checkPrefix(itemsPrefix, line), ", ")
	result := make([]int, len(parts))
	for i, item := range parts {
		result[i] = lib.ParseInt(item)
	}
	return result
}

const opPrefix = "  Operation: new = old "

func readOperation(line string) Op {
	parts := strings.Split(checkPrefix(opPrefix, line), " ")
	var op Op
	op.operator = parts[0]
	if parts[1] == "old" {
		op.useVal = true
	} else {
		op.value = lib.ParseInt(parts[1])
	}
	return op
}

type Op struct {
	operator string
	useVal   bool
	value    int
}

func (o Op) apply(v int) int {
	otherVal := o.value
	if o.useVal {
		otherVal = v
	}
	switch o.operator {
	case "+":
		return v + otherVal
	case "*":
		return v * otherVal
	}
	panic(fmt.Sprintf("unknown op %q", o.operator))
}

const (
	testPrefix  = "  Test: divisible by "
	truePrefix  = "    If true: throw to monkey "
	falsePrefix = "    If false: throw to monkey "
)

func readNext(lines []string) Next {
	return Next{
		divisor: lib.ParseInt(checkPrefix(testPrefix, lines[0])),
		yes:     lib.ParseInt(checkPrefix(truePrefix, lines[1])),
		no:      lib.ParseInt(checkPrefix(falsePrefix, lines[2])),
	}
}

type Next struct {
	divisor, yes, no int
}

func (n Next) apply(v int) int {
	if v%n.divisor == 0 {
		return n.yes
	}
	return n.no
}
