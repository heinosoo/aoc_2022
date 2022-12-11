package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	monkeyInputs := lines.Split(func(a string) bool { return a == "" })

	var monkeys []*Monkey
	divBy := 1
	for monkeyInput := range monkeyInputs {
		monkey := parseMonkey(monkeyInput)
		divBy *= monkey.divBy
		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 10000; i++ {
		round2(monkeys, divBy)
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

func round2(monkeys []*Monkey, divBy int) {
	for _, monkey := range monkeys {
		for !monkey.items.Empty() {
			monkey.N++
			oldWorry, _ := monkey.items.Dequeue()
			worry := monkey.op(oldWorry.(int)) % divBy
			monkeys[monkey.next(worry)].items.Enqueue(worry)
		}
	}
}
