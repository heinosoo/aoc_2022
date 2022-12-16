package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	V := readValves(lines)
	open := make(map[string]bool)
	for label, v := range V {
		if v.rate == 0 {
			open[label] = true
		}
	}

	cache := make(map[string]int)
	maxRelease := tick("AA", "AA", open, V, 0, 26, 26, cache)
	Log(maxRelease)
}
