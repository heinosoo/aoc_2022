package main

import (
	"bufio"
	"os"
	"strconv"

	aoc_utils "github.com/heinosoo/aoc_2022"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	biggestMeal := 0
	for n := parseElf(scanner); n > 0; n = parseElf(scanner) {
		biggestMeal = aoc_utils.Max(n, biggestMeal)
	}
	return biggestMeal
}

func parseElf(scanner *bufio.Scanner) int {
	calories := 0
	for scanner.Scan() {
		n, e := strconv.Atoi(scanner.Text())
		if e != nil {
			break
		}
		calories += n
	}
	return calories
}
