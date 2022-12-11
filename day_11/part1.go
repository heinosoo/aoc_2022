package main

import (
	"strconv"
	"strings"

	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	monkeyInputs := lines.Split(func(a string) bool { return a == "" })

	var monkeys []*Monkey
	for monkeyInput := range monkeyInputs {
		monkeys = append(monkeys, parseMonkey(monkeyInput))
	}

	for i := 0; i < 20; i++ {
		round(monkeys)
	}

	a, b := 0, 0
	for _, monkey := range monkeys {
		n := monkey.N
		if n > a {
			a, b = n, a
		} else if n > b {
			b = n
		}
	}
	Log("Monkey business: ", a*b)

}

func round(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		for !monkey.items.Empty() {
			monkey.N++
			oldWorry, _ := monkey.items.Dequeue()
			worry := monkey.op(oldWorry.(int)) / 3
			monkeys[monkey.next(worry)].items.Enqueue(worry)
		}
	}
}

type Monkey struct {
	items *llq.Queue
	op    func(int) int
	next  func(int) int
	divBy int
	N     int
}

func parseMonkey(input Channel[string]) *Monkey {
	<-input // drop monkey number
	monkey := Monkey{items: llq.New()}

	items := strings.Split(<-input, " ")[4:]
	for _, item := range items {
		if item[len(item)-1] == ',' {
			item = item[:len(item)-1]
		}
		worry, _ := strconv.Atoi(item)
		monkey.items.Enqueue(worry)
	}

	opValues := strings.Split(<-input, " ")[5:]
	monkey.op = func(old int) (new int) {
		a, e := strconv.Atoi(opValues[0])
		if e != nil {
			a = old
		}
		b, e := strconv.Atoi(opValues[2])
		if e != nil {
			b = old
		}
		switch opValues[1] {
		case "+":
			return a + b
		case "*":
			return a * b
		default:
			panic("Unexpected operator.")
		}
	}

	monkey.divBy, _ = strconv.Atoi(strings.Split(<-input, " ")[5])
	trueMonkey, _ := strconv.Atoi(strings.Split(<-input, " ")[9])
	falseMonkey, _ := strconv.Atoi(strings.Split(<-input, " ")[9])

	monkey.next = func(a int) int {
		if a%monkey.divBy == 0 {
			return trueMonkey
		} else {
			return falseMonkey
		}
	}
	return &monkey
}
