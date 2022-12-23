package main

import (
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

type monkeyT struct {
	name        string
	val         int
	op          byte
	human       bool
	left, right *monkeyT
}

func part1(lines Channel[string]) {
	monkeys := parseMonkeys(lines)
	Log(monkeys["root"].calculate())
}

func (monkey *monkeyT) calculate() int {
	if monkey.val == 0 {
		left, right := monkey.left.calculate(), monkey.right.calculate()
		switch monkey.op {
		case '+':
			monkey.val = left + right
		case '-':
			monkey.val = left - right
		case '*':
			monkey.val = left * right
		case '/':
			monkey.val = left / right
		}
	}
	return monkey.val
}

func parseMonkeys(lines Channel[string]) (monkeys map[string]*monkeyT) {
	monkeys = make(map[string]*monkeyT)
	opMap := make(map[string][2]string)
	for line := range lines {
		split := strings.Split(line, " ")
		name, split := split[0][:4], split[1:]
		if len(split) == 1 {
			val, _ := strconv.Atoi(split[0])
			monkeys[name] = &monkeyT{name, val, 0, false, nil, nil}
		} else {
			opMap[name] = [2]string{split[0], split[2]}
			monkeys[name] = &monkeyT{name, 0, split[1][0], false, nil, nil}
		}
	}
	for name, M := range opMap {
		(monkeys[name]).left = monkeys[M[0]]
		(monkeys[name]).right = monkeys[M[1]]
	}
	return
}
