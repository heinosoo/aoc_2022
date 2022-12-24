package main

import (
	"fmt"

	. "github.com/heinosoo/aoc_2022"
)

type loc [2]int

type elvesT map[loc]bool

func part1(lines Channel[string]) {
	elves := parseGrove(lines)
	for i := 0; i < 10; i++ {
		elves.round(i)
	}
	elves.print()
}

var dN [3]loc = [3]loc{{-1, -1}, {0, -1}, {1, -1}}
var dS [3]loc = [3]loc{{-1, 1}, {0, 1}, {1, 1}}
var dW [3]loc = [3]loc{{-1, -1}, {-1, 0}, {-1, 1}}
var dE [3]loc = [3]loc{{1, -1}, {1, 0}, {1, 1}}
var dAll [4][3]loc = [4][3]loc{dN, dS, dW, dE}
var dSurr [8]loc = [8]loc{{-1, -1}, {0, -1}, {1, -1}, {-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}}

func (elves *elvesT) round(start int) bool {
	proposals := make(map[loc][]loc)
	for elf := range *elves {
		if !elves.empty(elf) {
			for i := 0; i < 4; i++ {
				D := dAll[(i+start)%4]
				if elves.noElf(elf, D) {
					proposal := elf.add(D[1])
					proposals[proposal] = append(proposals[proposal], elf)
					break
				}
			}
		}
	}
	if len(proposals) == 0 {
		return true
	}
	for dest, src := range proposals {
		if len(src) == 1 {
			(*elves)[dest] = true
			delete(*elves, src[0])
		}
	}
	return false
}

func (elves *elvesT) noElf(elf loc, D [3]loc) bool {
	for _, d := range D {
		if (*elves)[elf.add(d)] {
			return false
		}
	}
	return true
}

func (elves *elvesT) empty(elf loc) bool {
	for _, d := range dSurr {
		if (*elves)[elf.add(d)] {
			return false
		}
	}
	return true
}

func (a loc) add(b loc) loc {
	return loc{a[0] + b[0], a[1] + b[1]}
}

func (elves *elvesT) print() {
	x1, x2, y1, y2 := 0, 0, 0, 0
	for L, _ := range *elves {
		x1, x2 = Min(x1, L[0]), Max(x2, L[0])
		y1, y2 = Min(y1, L[1]), Max(y2, L[1])
	}
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if (*elves)[loc{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		Log()
	}
	Log((x2-x1+1)*(y2-y1+1) - len(*elves))
}

func parseGrove(lines Channel[string]) (elves elvesT) {
	elves = make(elvesT)
	y := 0
	for line := range lines {
		for x, c := range line {
			if c == '#' {
				elves[loc{x, y}] = true
			}
		}
		y++
	}
	return
}
