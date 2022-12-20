package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	numbers, zero := parseFile(lines, 811589153)
	for i := 0; i < 10; i++ {
		for _, num := range numbers {
			num.shift()
		}
	}
	Log(zero.get(1000).val + zero.get(2000).val + zero.get(3000).val)
}
