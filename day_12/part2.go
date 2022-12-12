package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	minSteps := 239582395876
	heights, starts := readMap2(lines)
	for _, start := range starts {
		var visited [][]bool
		for i := 0; i < len(heights); i++ {
			visited = append(visited, make([]bool, len(heights[0])))
		}
		visited[start[0]][start[1]] = true
		edge := [][2]int{start}
		for i := 1; ; i++ {
			if len(edge) == 0 {
				break
			}
			if step(heights, visited, &edge) {
				minSteps = Min(minSteps, i)
				break
			}
		}
	}
	Log(minSteps)
}

func readMap2(lines Channel[string]) (heights [][]byte, starts [][2]int) {
	i := 0
	var start [2]int
	for line := range lines {
		i++
		for j, h := range line {
			if h == 'S' {
				start = [2]int{i, j + 1}
				starts = append(starts, [2]int{i, j + 1})
			} else if h == 'a' {
				starts = append(starts, [2]int{i, j + 1})
			}
		}

		heights = append(heights, append(append([]byte{0}, line...), 0))
	}
	heights = append([][]byte{make([]byte, len(heights[0]))}, heights...)
	heights = append(heights, make([]byte, len(heights[0])))
	heights[start[0]][start[1]] = 'a'
	return
}
