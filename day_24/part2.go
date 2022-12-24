package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	end, walk := parseMap(lines)
	start := loc{0, 0}
	steps := walk(start, end, 0)
	steps = walk(end, start, steps+1)
	steps = walk(start, end, steps+1)
	Log(steps)
}
