package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func main() {
	lines1, lines2 := make(chan string, 10), make(chan string, 10)
	go ReadFile(GetFilename(), lines1)
	go ReadFile(GetFilename(), lines2)

	WithTiming("Part 1 finished:", func() { part1(lines1) })
	Log()
	WithTiming("Part 2 finished:", func() { part2(lines2) })
}
