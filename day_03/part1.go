package main

import (
	"bufio"
	"os"
	"strings"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	commonItem("abcdefgd")

	s := 0
	for scanner.Scan() {
		s += priority(int(commonItem(scanner.Text())))
	}

	return s
}

func priority(item int) int {
	if int('A') <= item && item <= int('Z') {
		return item - 'A' + 27
	} else if int('a') <= item && item <= int('z') {
		return item - 'a' + 1
	} else {
		panic("Weird item.")
	}
}

func commonItem(rucksack string) rune {
	firstCompartment := rucksack[:len(rucksack)/2]
	secondCompartment := rucksack[len(rucksack)/2:]
	for _, item := range firstCompartment {
		if strings.ContainsRune(secondCompartment, item) {
			return item
		}
	}
	panic("No common item found.")
}
