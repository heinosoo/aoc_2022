package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	monkeys := parseMonkeys(lines)
	monkeys["root"].op = '='
	monkeys["root"].calculateNonhuman()

	Log(monkeys["root"].humanVal())
}

func (monkey *monkeyT) humanVal() int {
	Log(monkey.name, monkey.left.name, monkey.right.name)
	Log(string(monkey.op), monkey.val, monkey.left.val, monkey.right.val)

	if monkey.left.name == "humn" {
		return monkey.right.val + monkey.val
	} else if monkey.right.name == "humn" {
		return monkey.left.val + monkey.val
	} else if monkey.left.val == 0 {
		switch monkey.op {
		case '+':
			monkey.left.val = monkey.val - monkey.right.val
		case '*':
			monkey.left.val = monkey.val / monkey.right.val
		case '-':
			monkey.left.val = monkey.val + monkey.right.val
		case '/':
			monkey.left.val = monkey.val * monkey.right.val
		case '=':
			monkey.left.val = monkey.right.val
		}
		return monkey.left.humanVal()
	} else if monkey.right.val == 0 {
		switch monkey.op {
		case '+':
			monkey.right.val = monkey.val - monkey.left.val
		case '*':
			monkey.right.val = monkey.val / monkey.left.val
		case '-':
			monkey.right.val = monkey.left.val - monkey.val
		case '/':
			monkey.right.val = monkey.left.val / monkey.val
		case '=':
			monkey.right.val = monkey.left.val
		}
		return monkey.right.humanVal()
	}
	panic("Shouldn't get here.")
}

func (monkey *monkeyT) calculateNonhuman() (int, bool) {
	if monkey.val != 0 {
		return monkey.val, monkey.name == "humn"
	}
	if monkey.val == 0 {
		left, LH := monkey.left.calculateNonhuman()
		right, RH := monkey.right.calculateNonhuman()
		if LH || RH {
			return 0, true
		}
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
	return monkey.val, false
}
