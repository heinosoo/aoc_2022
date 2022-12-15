package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	S, B := readMap(lines)
	tuning := make(chan int)
	for s, r := range S {
		go checkAround(s, r, tuning, S, B)
	}
	Log("Waiting for tuning")
	Log(<-tuning)
}

func checkAround(P point, r int, tuning chan int, S sensors, B beacons) {
	pointChannel := circle(P, r)
	for P := range pointChannel {
		if !B[P] && S[P] == 0 {
			go func(P point) {
				for s, r := range S {
					if distance(s, P) <= r {
						return
					}
				}
				tuning <- P[0]*4_000_000 + P[1]
			}(P)
		}
	}
}

func circle(A point, r int) Channel[point] {
	channel := NewChannel[point](1)
	min, max := 0, 4_000_000 // 20
	go func() {
		for dx := -r - 1; dx <= r+1; dx++ {
			dy := r + 1 - dx
			x, y1, y2 := A[0]+dx, A[1]+dy, A[1]-dy
			if x >= min && x <= max && y1 >= min && y1 <= max {
				channel <- point{x, y1}
			}
			if x >= min && x <= max && y2 >= min && y2 <= max {
				channel <- point{x, y2}
			}
		}
		close(channel)
	}()
	return channel
}
