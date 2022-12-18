package main

import (
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

type cubeT [3]int
type dropletT struct {
	S int
	M map[cubeT]bool
}

// input_test - 64
func part1(lines Channel[string]) {
	droplet := dropletT{0, make(map[cubeT]bool)}
	for line := range lines {
		var cube cubeT
		for i, nString := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(nString)
			cube[i] = n
		}
		droplet.addCube(cube)
	}
	Log(droplet.S)
}

func (droplet *dropletT) addCube(cube cubeT) {
	droplet.M[cube] = true
	droplet.S += 6
	for _, cube := range cube.around() {
		if droplet.M[cube] {
			droplet.S -= 2
		}
	}
}

func (cube *cubeT) around() (cubes [6]cubeT) {
	D := [6][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	for i, d := range D {
		for j := 0; j < 3; j++ {
			cubes[i][j] = cube[j] + d[j]
		}
	}
	return
}
