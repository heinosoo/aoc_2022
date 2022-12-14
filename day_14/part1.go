package main

import (
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	cave, maxY := createMap(lines)
	i := 0
	for more := true; more; i++ {
		_, more = addSand(cave, [2]int{500, 0}, maxY)
	}
	Log(i - 1)
}

func addSand(cave map[[2]int]bool, loc [2]int, maxY int) (last [2]int, more bool) {
	for loc[1] <= maxY {
		next := next(cave, loc)
		if next == loc {
			(cave)[loc] = true
			return loc, true
		}
		loc = next
	}
	return loc, false
}

func next(cave map[[2]int]bool, sand [2]int) [2]int {
	switch {
	case !(cave)[[2]int{sand[0], sand[1] + 1}]:
		return [2]int{sand[0], sand[1] + 1}
	case !(cave)[[2]int{sand[0] - 1, sand[1] + 1}]:
		return [2]int{sand[0] - 1, sand[1] + 1}
	case !(cave)[[2]int{sand[0] + 1, sand[1] + 1}]:
		return [2]int{sand[0] + 1, sand[1] + 1}
	default:
		return sand
	}
}

func createMap(lines Channel[string]) (cave map[[2]int]bool, maxY int) {
	cave = make(map[[2]int]bool)
	for line := range lines {
		points := getPoints(line)
		P1 := points[0]
		maxY = Max(maxY, P1[1])
		for _, P2 := range points[1:] {
			maxY = Max(maxY, P2[1])
			switch {
			case P1[0] < P2[0]:
				for x := P1[0]; x <= P2[0]; x++ {
					cave[[2]int{x, P1[1]}] = true
				}
			case P1[0] > P2[0]:
				for x := P1[0]; x >= P2[0]; x-- {
					cave[[2]int{x, P1[1]}] = true
				}
			case P1[1] < P2[1]:
				for y := P1[1]; y <= P2[1]; y++ {
					cave[[2]int{P1[0], y}] = true
				}
			case P1[1] > P2[1]:
				for y := P1[1]; y >= P2[1]; y-- {
					cave[[2]int{P1[0], y}] = true
				}
			}
			P1 = P2
		}
	}
	return
}

func getPoints(line string) (points [][2]int) {
	for _, pointString := range strings.Split(line, " -> ") {
		point := strings.Split(pointString, ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		points = append(points, [2]int{x, y})
	}
	return
}
