package main

import (
	"bufio"
	"os"

	aoc_utils "github.com/heinosoo/aoc_2022"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	biggestMeals := make([]int, 3)
	for n := parseElf(scanner); n > 0; n = parseElf(scanner) {
		smallestBiggestMeal, i := aoc_utils.MinOf(biggestMeals...)
		biggestMeals[i] = aoc_utils.Max(smallestBiggestMeal, n)
	}

	return aoc_utils.Sum(biggestMeals...)
}
