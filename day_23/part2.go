package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	elves := parseGrove(lines)
	for i := 0; i < 100000; i++ {
		if elves.round(i) {
			Log(i + 1)
			break
		}
	}
}
