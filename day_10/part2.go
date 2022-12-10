package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	buffer := make(chan [2]int, 10)
	go parseCommands(scanner, buffer)

	line := ""
	for {
		state, more := <-buffer
		if !more {
			break
		} else if (state[0])%40 == 0 {
			line += pixel(state)
			fmt.Println(line)
			line = ""
		} else {
			line += pixel(state)
		}
	}

	return 1
}

func pixel(state [2]int) string {
	crt, sprite := (state[0]-1)%40, state[1]
	if -1 <= crt-sprite && crt-sprite <= 1 {
		return "#"
	} else {
		return "."
	}
}
