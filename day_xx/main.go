package main

import (
	"fmt"

	aoc_utils "github.com/heinosoo/aoc_2022"
)

func main() {
	filename := aoc_utils.GetFilename()
	fmt.Println("Part 1: ", part1(filename))
	fmt.Println("Part 2: ", part2(filename))
}
