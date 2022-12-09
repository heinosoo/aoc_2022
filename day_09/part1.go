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

	visited := map[[2]int]bool{{0, 0}: true}
	head, tail := [2]int{0, 0}, [2]int{0, 0}
	for scanner.Scan() {
		dir, n := parseLine(scanner.Text())
		for i := 0; i < n; i++ {
			update(&head, &tail, dir)
			visited[tail] = true
		}
	}

	return len(visited)
}

func parseLine(line string) (dir byte, n int) {
	sp := strings.Split(line, " ")
	dir = sp[0][0]
	n, _ = strconv.Atoi(sp[1])
	return
}

var DIR = map[byte][2]int{'U': {0, 1}, 'D': {0, -1}, 'R': {1, 0}, 'L': {-1, 0}}

func move(end *[2]int, dir byte, opposite bool) {
	if opposite {
		end[0] -= DIR[dir][0]
		end[1] -= DIR[dir][1]
	} else {
		end[0] += DIR[dir][0]
		end[1] += DIR[dir][1]
	}
}

func update(head, tail *[2]int, dir byte) {
	move(head, dir, false)
	diff := [2]int{head[0] - tail[0], head[1] - tail[1]}
	d := diff[0]*diff[0] + diff[1]*diff[1]
	switch {
	case d <= 2:
		return
	case d == 4:
		move(tail, dir, false)
	case d == 5:
		tail[0], tail[1] = head[0], head[1]
		move(tail, dir, true)
	}
}
