package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	buffer := make(chan [2]int, 10)
	go parseCommands(scanner, buffer)

	answer := 0
	for {
		state, more := <-buffer
		if !more {
			break
		} else if more && ((state[0]+20)%40 == 0) {
			answer += state[0] * state[1]
		}
	}

	return answer
}

func parseCommands(scanner *bufio.Scanner, buffer chan [2]int) {
	state := [2]int{0, 1}
	state = noop(buffer, state)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "noop":
			state = noop(buffer, state)
		case "addx":
			a, _ := strconv.Atoi(line[1])
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
