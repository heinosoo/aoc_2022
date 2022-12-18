package main

import (
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

// input_test - 58
func part2(lines Channel[string]) {
	droplet := dropletT{0, make(map[cubeT]bool)}
	limits := [3][2]int{}
	for line := range lines {
		var cube cubeT
		for i, nString := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(nString)
			cube[i] = n
			limits[i][0] = Min(limits[i][0], n-1)
			limits[i][1] = Max(limits[i][1], n+1)
		}
		droplet.addCube(cube)
	}

	water := dropletT{0, make(map[cubeT]bool)}
	water.fill(droplet, limits)
	x, y, z := limits[0][1]-limits[0][0]+1, limits[1][1]-limits[1][0]+1, limits[2][1]-limits[2][0]+1

	Log(water.S - 2*(x*y+x*z+y*z))
}

func (surrounding *dropletT) fill(droplet dropletT, limits [3][2]int) {
	edge := []cubeT{{limits[0][0], limits[1][0], limits[2][0]}}
	for len(edge) != 0 {
		newEdge := []cubeT{}
		for _, cube := range edge {
			for _, cube := range cube.around() {
				if !surrounding.M[cube] && !droplet.M[cube] && !cube.offLimits(limits) {
					surrounding.addCube(cube)
					newEdge = append(newEdge, cube)
				}
			}
		}
		edge = newEdge
	}
}

func (cube *cubeT) offLimits(limits [3][2]int) bool {
	for i := 0; i < 3; i++ {
		if cube[i] < limits[i][0] || cube[i] > limits[i][1] {
			return true
		}
	}
	return false
}
