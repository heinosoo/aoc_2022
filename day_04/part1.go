package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	s := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		sectionsA := parseElf(line[0])
		sectionsB := parseElf(line[1])
		if fullyContains(sectionsA, sectionsB) || fullyContains(sectionsB, sectionsA) {
			s += 1
		}
	}

	return s
}

func parseElf(elf string) (sections [2]int) {
	sectionsString := strings.Split(elf, "-")
	sections[0], _ = strconv.Atoi(sectionsString[0])
	sections[1], _ = strconv.Atoi(sectionsString[1])
	return
}

func fullyContains(elfA [2]int, elfB [2]int) bool {
	return elfA[0] <= elfB[0] && elfB[1] <= elfA[1]
}
