package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	V := readValves(lines)
	maxRelease := runTickers(26, 26, V)
	Log(maxRelease)
}
