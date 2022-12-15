package main

import (
	"regexp"
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	S, B := readMap(lines)
	l := 2_000_000 // 10

	covered := make(map[int]bool)
	for s, r := range S {
		xr := r - Abs(s[1]-l)
		if xr < 0 {
			continue
		}
		for x := s[0] - xr; x <= s[0]+xr; x++ {
			covered[x] = true
		}
	}

	noBeacons := len(covered)
	for b := range B {
		if b[1] == l {
			noBeacons--
		}
	}
	Log(noBeacons)
}

type point [2]int
type beacons map[point]bool
type sensors map[point]int

func readMap(lines Channel[string]) (S sensors, B beacons) {
	S, B = make(sensors), make(beacons)
	for line := range lines {
		r := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`).FindStringSubmatch(line)
		sx, _ := strconv.Atoi(r[1])
		sy, _ := strconv.Atoi(r[2])
		bx, _ := strconv.Atoi(r[3])
		by, _ := strconv.Atoi(r[4])
		s, b := point{sx, sy}, point{bx, by}
		S[s], B[b] = distance(s, b), true
	}
	return
}

func distance(a [2]int, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}
