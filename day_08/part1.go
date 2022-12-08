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

	forest := readForest(scanner)
	hidden := scanHidden(forest)

	s := 0
	aoc_utils.OperateOnMatrix(&hidden, func(a bool) {
		if !a {
			s++
		}
	})

	return s
}

func readForest(scanner *bufio.Scanner) (forest [][]int) {
	for scanner.Scan() {
		var treeLine []int
		for _, tree := range scanner.Bytes() {
			height, _ := strconv.Atoi(string(tree))
			treeLine = append(treeLine, height)
		}
		forest = append(forest, treeLine)
	}
	return
}

func scanHidden(forest [][]int) (hidden [][]bool) {
	M, N := len(forest), len(forest[0])
	hidden = aoc_utils.InitializeMatrix(false, M, N)

	for m := 1; m < M-1; m++ {
		tallest := 0
		for n := 1; n < N-1; n++ {
			tallest = aoc_utils.Max(tallest, forest[m][n-1])
			hidden[m][n] = forest[m][n] <= tallest
		}
		tallest = 0
		for n := N - 2; n > 0; n-- {
			tallest = aoc_utils.Max(tallest, forest[m][n+1])
			hidden[m][n] = hidden[m][n] && (forest[m][n] <= tallest)
		}
	}
	for m := 1; m < N-1; m++ {
		tallest := 0
		for n := 1; n < M-1; n++ {
			tallest = aoc_utils.Max(tallest, forest[n-1][m])
			hidden[n][m] = hidden[n][m] && (forest[n][m] <= tallest)
		}
		tallest = 0
		for n := M - 2; n > 0; n-- {
			tallest = aoc_utils.Max(tallest, forest[n+1][m])
			hidden[n][m] = hidden[n][m] && (forest[n][m] <= tallest)
		}
	}
	return

}
