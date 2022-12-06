package main

import (
	"bufio"
	"os"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bytes := scanner.Bytes()

	unique := uniqueMaker(14)
	for i, b := range bytes {
		if unique(b) {
			return (i + 1)
		}
	}

	panic("Didn't find marker.")
}
