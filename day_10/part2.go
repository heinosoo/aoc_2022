package main

import (
	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	buffer := NewChannel[[2]int](10)
	go parseCommands(lines, buffer)

	line := ""
	for state := range buffer {
		if (state[0])%40 == 0 {
			line += pixel(state)
			Log(line)
			line = ""
		} else {
			line += pixel(state)
		}
	}
}

func pixel(state [2]int) string {
	crt, sprite := (state[0]-1)%40, state[1]
	if -1 <= crt-sprite && crt-sprite <= 1 {
		return "#"
	} else {
		return "."
	}
}
