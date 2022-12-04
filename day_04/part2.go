package main

import (
	"bufio"
	"os"
	"strings"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	s := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		sectionsA := parseElf(line[0])
		sectionsB := parseElf(line[1])
		if partiallyContains(sectionsA, sectionsB) || partiallyContains(sectionsB, sectionsA) {
			s += 1
		}
	}

	return s
}

func partiallyContains(elfA [2]int, elfB [2]int) bool {
	return (elfA[0] <= elfB[0] && elfB[0] <= elfA[1]) || (elfA[0] <= elfB[1] && elfB[1] <= elfA[1])
}
