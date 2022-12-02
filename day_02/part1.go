package main

import (
	"bufio"
	"os"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		elf, me := scanner.Text()[0], scanner.Text()[2]
		score += match(elf, shapeMap[me])
	}

	return score
}

var scoreMap = map[byte]int{'A': 1, 'B': 2, 'C': 3}
var beatMap = map[byte]byte{'A': 'C', 'B': 'A', 'C': 'B'}
var shapeMap = map[byte]byte{'X': 'A', 'Y': 'B', 'Z': 'C'}

func match(elf byte, me byte) int {
	score := scoreMap[me]
	if me == elf {
		score += 3
	} else if beatMap[me] == elf {
		score += 6
	}
	return score
}
