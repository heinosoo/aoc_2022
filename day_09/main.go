package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func main() {
	lines1, lines2 := make(chan string, 10), make(chan string, 10)
	go ReadFile(GetFilename(), lines1)
	go ReadFile(GetFilename(), lines2)

	part1(lines1)
	part2(lines2)
}
