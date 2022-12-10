package main

import (
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines chan string) {
	Log("Part 1:")

	forest := readForest(lines)
	hidden := scanHidden(forest)

	s := 0
	OperateOnMatrix(&hidden, func(a bool) {
		if !a {
			s++
		}
	})

	Log(s)
}

func readForest(lines chan string) (forest [][]int) {
	for line := range lines {
		var treeLine []int
		for _, tree := range line {
			height, _ := strconv.Atoi(string(tree))
			treeLine = append(treeLine, height)
		}
		forest = append(forest, treeLine)
	}
	return
}

func scanHidden(forest [][]int) (hidden [][]bool) {
	M, N := len(forest), len(forest[0])
	hidden = InitializeMatrix(false, M, N)

	for m := 1; m < M-1; m++ {
		tallest := 0
		for n := 1; n < N-1; n++ {
			tallest = Max(tallest, forest[m][n-1])
			hidden[m][n] = forest[m][n] <= tallest
		}
		tallest = 0
		for n := N - 2; n > 0; n-- {
			tallest = Max(tallest, forest[m][n+1])
			hidden[m][n] = hidden[m][n] && (forest[m][n] <= tallest)
		}
	}
	for m := 1; m < N-1; m++ {
		tallest := 0
		for n := 1; n < M-1; n++ {
			tallest = Max(tallest, forest[n-1][m])
			hidden[n][m] = hidden[n][m] && (forest[n][m] <= tallest)
		}
		tallest = 0
		for n := M - 2; n > 0; n-- {
			tallest = Max(tallest, forest[n+1][m])
			hidden[n][m] = hidden[n][m] && (forest[n][m] <= tallest)
		}
	}
	return

}
