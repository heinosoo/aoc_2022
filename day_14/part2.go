package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	cave, maxY := createMap(lines)
	i, last := 0, [2]int{0, 0}
	for ; last != [2]int{500, 0}; i++ {
		last, _ = addSand(cave, [2]int{500, 0}, maxY)
		cave[last] = true
	}
	Log(i)
}
