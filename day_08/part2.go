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

	forest := readForest(scanner)
	score := scenicScore2(forest)

	max := 0
	aoc_utils.OperateOnMatrix(&score, func(a int) {
		max = aoc_utils.Max(max, a)
	})

	return max
}

func scenicScore2(forest [][]int) (score [][]int) {
	M, N := len(forest), len(forest[0])
	score = aoc_utils.InitializeMatrix(1, M, N)
	for H := range [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		for m := 0; m < M; m++ {
			score[m][0], score[m][N-1] = 0, 0
			distance := 1
			for n := 1; n < N; n++ {
				tree := forest[m][n]
				switch {
				case tree == H:
					score[m][n] *= distance
					distance = 1
				case tree > H:
					distance = 1
				case tree < H:
					distance++
				}
			}
			distance = 1
			for n := N - 2; n >= 0; n-- {
				tree := forest[m][n]
				switch {
				case tree == H:
					score[m][n] *= distance
					distance = 1
				case tree > H:
					distance = 1
				case tree < H:
					distance++
				}
			}
		}
		for m := 0; m < N; m++ {
			score[0][m], score[M-1][m] = 0, 0
			distance := 1
			for n := 1; n < M; n++ {
				tree := forest[n][m]
				switch {
				case tree == H:
					score[n][m] *= distance
					distance = 1
				case tree > H:
					distance = 1
				case tree < H:
					distance++
				}
			}
			distance = 1
			for n := M - 2; n >= 0; n-- {
				tree := forest[n][m]
				switch {
				case tree == H:
					score[n][m] *= distance
					distance = 1
				case tree > H:
					distance = 1
				case tree < H:
					distance++
				}
			}
		}
	}
	return
}
