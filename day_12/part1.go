package main

import (
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	heights, visited, start := readMap(lines)
	edge := [][2]int{start}
	for i := 1; ; i++ {
		if step(heights, visited, &edge) {
			Log(i)
			break
		}
	}
}

func step(heights [][]byte, visited [][]bool, edge *[][2]int) bool {
	var newEdge [][2]int
	for _, s := range *edge {
		for _, e := range surrounding(s) {
			h, v := heights[e[0]][e[1]], visited[e[0]][e[1]]
			if v || h == 0 {
				continue
			} else if h == 'E' {
				if heights[s[0]][s[1]] >= 'z'-1 {
					return true
				}
			} else if heights[s[0]][s[1]] >= h-1 {
				visited[e[0]][e[1]] = true
				newEdge = append(newEdge, e)
			} else if h == 'E' {
			}
		}
	}
	*edge = newEdge
	return false
}

func surrounding(S [2]int) (points [][2]int) {
	for _, D := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		points = append(points, [2]int{S[0] + D[0], S[1] + D[1]})
	}
	return
}

func readMap(lines Channel[string]) (heights [][]byte, visited [][]bool, start [2]int) {
	i := 0
	for line := range lines {
		i++
		j := strings.Index(line, "S")
		if j != -1 {
			start = [2]int{i, j + 1}
		}
		heights = append(heights, append(append([]byte{0}, line...), 0))
	}
	heights = append([][]byte{make([]byte, len(heights[0]))}, heights...)
	heights = append(heights, make([]byte, len(heights[0])))

	for i := 0; i < len(heights); i++ {
		visited = append(visited, make([]bool, len(heights[0])))
	}
	visited[start[0]][start[0]] = true
	heights[start[0]][start[1]] = 'a'
	return
}
