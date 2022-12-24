package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	end, walk := parseMap(lines)
	start := loc{0, 0}
	Log(walk(start, end, 0))
}

var DIRMAP = map[rune]int{'>': 0, '<': 1, 'v': 2, '^': 3}
var DIR = [4]loc{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func parseMap(lines Channel[string]) (end loc, walk func(start, end loc, i int) int) {
	B := blizzards{make(map[loc]bool), make(map[loc]bool), make(map[loc]bool), make(map[loc]bool)}
	Nx, y := 0, 0
	for line := range lines {
		Nx = len(line) - 2
		for x, c := range line {
			switch c {
			case '#':
			case '.':
			default:
				B[DIRMAP[c]][loc{x - 1, y - 1}] = true
			}
		}
		y++
	}
	Ny := y - 2
	end = loc{Nx - 1, Ny - 1}

	isAllowed := func(i int, p loc) bool {
		if p[0] < 0 || p[0] >= Nx || p[1] < 0 || p[1] >= Ny {
			return false
		} else {
			p0 := loc{(p[0] - i%Nx + Nx) % Nx, p[1]}
			p1 := loc{(p[0] + i) % Nx, p[1]}
			p2 := loc{p[0], (p[1] - i%Ny + Ny) % Ny}
			p3 := loc{p[0], (p[1] + i) % Ny}
			return !B[0][p0] && !B[1][p1] && !B[2][p2] && !B[3][p3]
		}
	}

	walk = func(start, end loc, i int) int {
		edge := make(map[loc]bool)
		for ; ; i++ {
			if isAllowed(i, start) {
				edge[start] = true
			}
			newEdge := make(map[loc]bool)
			for p := range edge {
				if p == end {
					return i
				}
				if isAllowed(i, p) {
					newEdge[p] = true
				}
				for _, d := range DIR {
					pd := p.add(d)
					if isAllowed(i, pd) {
						newEdge[pd] = true
					}
				}
			}
			edge = newEdge
		}
	}

	return
}

type loc [2]int
type blizzards [4]map[loc]bool

func (a loc) add(b loc) (c loc) {
	c[0], c[1] = a[0]+b[0], a[1]+b[1]
	return
}
