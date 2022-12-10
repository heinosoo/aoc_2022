package main

import (
	"strconv"
	"strings"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines chan string) {
	Log("Part 1:")

	buffer := make(chan [2]int, 10)
	go parseCommands(lines, buffer)

	answer := 0
	for state := range buffer {
		if (state[0]+20)%40 == 0 {
			answer += state[0] * state[1]
		}
	}

	Log(answer)
}

func parseCommands(lines chan string, buffer chan [2]int) {
	state := [2]int{0, 1}
	state = noop(buffer, state)
	for line := range lines {
		lineSplit := strings.Split(line, " ")
		switch lineSplit[0] {
		case "noop":
			state = noop(buffer, state)
		case "addx":
			a, _ := strconv.Atoi(lineSplit[1])
			state = noop(buffer, state)
			state = addx(buffer, state, a)
		}
	}
	close(buffer)
}

func noop(buffer chan [2]int, state [2]int) [2]int {
	state[0]++
	buffer <- state
	return state
}

func addx(buffer chan [2]int, state [2]int, a int) [2]int {
	state[0]++
	state[1] += a
	buffer <- state
	return state
}
