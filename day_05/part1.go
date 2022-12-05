package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func part1(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	crates := crateLines(scanner)
	moves := make(chan [3]int, 2)
	go readMoves(scanner, moves)
	processMoves(crates, moves)

	message := ""
	for _, crate := range crates {
		top, _ := crate.Peek()
		message += string(top.(byte))
	}

	return message
}

func processMoves(crates []lls.Stack, moves chan [3]int) {
	for {
		move, more := <-moves
		if more {
			for i := 0; i < move[0]; i++ {
				crate, _ := crates[move[1]-1].Pop()
				crates[move[2]-1].Push(crate)
			}
		} else {
			return
		}
	}
}

func crateLines(scanner *bufio.Scanner) (crates []lls.Stack) {
	lines := lls.New()
	N := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			oldline, _ := lines.Pop()
			N = (len(oldline.(string)) + 1) / 4
			break
		}
		lines.Push(line)
	}

	for j := 0; j < N; j++ {
		crates = append(crates, *lls.New())
	}

	for line, more := lines.Pop(); more; line, more = lines.Pop() {
		for j := 0; j < N; j++ {
			a := line.(string)[1+4*j]
			if a != ' ' {
				crates[j].Push(a)
			}
		}
	}

	return
}

func readMoves(scanner *bufio.Scanner, moves chan [3]int) {
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(splitLine[1])
		b, _ := strconv.Atoi(splitLine[3])
		c, _ := strconv.Atoi(splitLine[5])
		moves <- [3]int{a, b, c}
	}
	close(moves)
}
