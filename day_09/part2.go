package main

import (
	"fmt"

	. "github.com/heinosoo/aoc_2022"
)

func part2(lines chan string) {
	Log("Part 2:")
	visited := map[[2]int]bool{{0, 0}: true}
	var knots [10][2]int
	for line := range lines {
		dir, n := parseLine(line)
		for i := 0; i < n; i++ {
			knots = update2(knots, dir)
			visited[knots[9]] = true
			// printPicture(knots, visited)
		}

	}
	// printPicture(knots, visited)
	Log(len(visited))
}

func update2(knots [10][2]int, dir byte) [10][2]int {
	knots[0] = move2(knots[0], dir, false)
	for i := range knots[1:] {
		diff := [2]int{knots[i][0] - knots[i+1][0], knots[i][1] - knots[i+1][1]}
		d := diff[0]*diff[0] + diff[1]*diff[1]
		switch {
		case d <= 2:
		case d == 4 && diff[0] == 0:
			knots[i+1] = move2(knots[i+1], 'U', diff[1] < 0)
		case d == 4 && diff[1] == 0:
			knots[i+1] = move2(knots[i+1], 'R', diff[0] < 0)
		case d >= 5:
			if diff[0] > 0 {
				knots[i+1] = move2(knots[i+1], 'R', false)
			} else {
				knots[i+1] = move2(knots[i+1], 'L', false)
			}
			if diff[1] > 0 {
				knots[i+1] = move2(knots[i+1], 'U', false)
			} else {
				knots[i+1] = move2(knots[i+1], 'D', false)
			}
		}
	}
	return knots
}

func move2(end [2]int, dir byte, opposite bool) [2]int {
	if opposite {
		end[0] -= DIR[dir][0]
		end[1] -= DIR[dir][1]
	} else {
		end[0] += DIR[dir][0]
		end[1] += DIR[dir][1]
	}
	return end
}

func printPicture(knots [10][2]int, visited map[[2]int]bool) {
	const size int = 15
	var picture [size*2 + 1][size*2 + 1]string
	picture[0][0] = "s"
	x, y := func(k [2]int) int { return (size - k[1]) }, func(k [2]int) int { return (size + k[0]) }

	for i, line := range picture {
		for j := range line {
			picture[i][j] = " "
		}
	}

	for k := range visited {
		picture[x(k)][y(k)] = "-"
	}

	for i, k := range knots {
		if i == 0 {
			picture[x(k)][y(k)] = "H"
		}
		if picture[x(k)][y(k)] == " " || picture[x(k)][y(k)] == "-" {
			picture[x(k)][y(k)] = fmt.Sprint(i)
		}

	}
	for _, line := range picture {
		Log(line)
	}
	Log()
	WaitForInput()
}
