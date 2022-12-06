package main

import (
	"bufio"
	"os"

	aoc_utils "github.com/heinosoo/aoc_2022"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bytes := scanner.Bytes()

	unique := uniqueMaker(4)
	for i, b := range bytes {
		if unique(b) {
			return (i + 1)
		}
	}

	panic("Didn't find marker.")
}

func uniqueMaker(N int) func(byte) bool {
	i := 0
	buffer := make([]byte, N-1)
	maskForward := N
	return func(new byte) bool {
		for n := N - 2; n >= 0; n-- {
			if buffer[(i+n)%(N-1)] == new {
				maskForward = aoc_utils.Max(maskForward-1, n)
				buffer[i%(N-1)] = new
				i++
				return false
			}
		}
		maskForward--
		buffer[i%(N-1)] = new
		i++
		return maskForward <= -1
	}
}
