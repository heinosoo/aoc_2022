package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func main() {
	lines1 := NewChannel[string](10)
	go ReadFile(GetFilename(), lines1)

	WithTiming("Part 1 finished:", func() { part1(lines1) })
}
