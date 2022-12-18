package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	N := 1_000_000_000_000
	boulders := parseBouldersFile("boulders")
	moves := parseMoves(<-lines)

	Log(initShaft(moves, boulders)(N))
}
