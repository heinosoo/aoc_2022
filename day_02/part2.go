package main

import (
	"bufio"
	"os"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		elf, strategy := scanner.Text()[0], scanner.Text()[2]
		score += match2(elf, strategy)
	}

	return score
}

// var scoreMap = map[byte]int{'A': 1, 'B': 2, 'C': 3}
// var beatMap = map[byte]byte{'A': 'C', 'B': 'A', 'C': 'B'}
var loseMap = map[byte]byte{'A': 'B', 'B': 'C', 'C': 'A'}
var drawMap = map[byte]byte{'A': 'A', 'B': 'B', 'C': 'C'}
var strategyMap = map[byte]map[byte]byte{'X': beatMap, 'Y': drawMap, 'Z': loseMap}

func match2(elf byte, strategy byte) (score int) {
	me := strategyMap[strategy][elf]
	score = scoreMap[me]
	if me == elf {
		score += 3
	} else if beatMap[me] == elf {
		score += 6
	}
	return
}
