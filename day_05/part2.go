package main

import (
	"bufio"
	"os"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func part2(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	crates := crateLines(scanner)
	moves := make(chan [3]int, 2)
	go readMoves(scanner, moves)
	processMoves2(crates, moves)

	message := ""
	for _, crate := range crates {
		top, _ := crate.Peek()
		message += string(top.(byte))
	}

	return message
}

func processMoves2(crates []lls.Stack, moves chan [3]int) {
	for {
		move, more := <-moves
		if more {
			stack := lls.New()
			for i := 0; i < move[0]; i++ {
				crate, _ := crates[move[1]-1].Pop()
				stack.Push(crate)
			}
			for crate, more := stack.Pop(); more; crate, more = stack.Pop() {
				crates[move[2]-1].Push(crate)
			}
		} else {
			return
		}
	}
}
