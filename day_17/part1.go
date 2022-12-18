package main

import (
	"fmt"

	. "github.com/heinosoo/aoc_2022"
)

type shaftT = [1000]int16
type boulderT = [4]int16

func part1(lines Channel[string]) {
	boulders := parseBouldersFile("boulders")
	moves := parseMoves(<-lines)

	Log(initShaft(moves, boulders)(2022))
}

func initShaft(moves []int, boulders []boulderT) (addN func(int) int) {
	var shaft shaftT
	shaft[0] = 1<<9 - 1
	walls := int16(1 | 1<<8)
	H, T, Y, X := 0, 0, 0, 3
	Bi, Bn := 0, len(boulders)
	Mi, Mn := 0, len(moves)

	settleBoulder := func() {
		for i, Brow := range boulders[Bi] {
			y := (Y + i + 1000) % 1000
			shaft[y] = shaft[y] | (Brow << (X + 1))
		}
		Bi = (Bi + 1) % Bn
		for i := 0; i < 10; i++ {
			if shaft[(T+4-i)%1000]-walls != 0 {
				T = (T + 4 - i) % 1000
				H += 4 - i
				return
			}
		}
		panic("Tower lowered somehow?")
	}
	wouldFit := func(dx, dy int) bool {
		for i, Brow := range boulders[Bi] {
			y := (Y + i + dy + 1000) % 1000
			if shaft[y]&(Brow<<(X+dx+1)) != 0 {
				return false
			}
		}
		return true
	}
	fall := func() (blocked bool) {
		if wouldFit(0, -1) {
			Y = (Y - 1 + 1000) % 1000
			return false
		} else {
			return true
		}
	}
	push := func() {
		if wouldFit(moves[Mi], 0) {
			X += moves[Mi]
		}
		Mi = (Mi + 1) % Mn
	}
	cacheKey := func() (key int64) {
		for i := 0; i < 40; i++ {
			key += int64(shaft[(T-i+1000)%1000]) << i * 9
		}
		key = key<<2 + int64(Mi)
		key = key<<2 + int64(Bi)

		return
	}
	addBoulder := func() {
		X, Y = 2, T+4

		for i := 1; i <= 7; i++ {
			shaft[(T+i)%1000] = walls
		}
		for {
			push()
			if fall() {
				settleBoulder()
				return
			}
		}
	}
	addN = func(N int) int {
		cache := make(map[int64][2]int)
		for i := 1; i <= N; i++ {
			addBoulder()
			cacheVal := cache[cacheKey()]
			lastH, j := cacheVal[0], cacheVal[1]
			if lastH == 0 {
				cache[cacheKey()] = [2]int{H, i}
			} else {
				skip := (N - i) / (i - j)
				H += (H - lastH) * skip
				i += (i - j) * skip
			}
		}
		return H
	}

	return
}

func parseBouldersFile(filename string) (boulders []boulderT) {
	lines := NewChannel[string](0)
	go ReadFile(filename, lines)

	current, y := boulderT{0, 0, 0, 0}, 0
	for line := range lines {
		if line == "" {
			boulders = append(boulders, current)
			current, y = boulderT{}, 0
			continue
		}
		var row int16
		for x, c := range line {
			if c == '#' {
				row = row | 1<<x //(len(line)-1-x)
			}
		}
		current[y] = row
		y++
	}
	boulders = append(boulders, current)
	return
}

func parseMoves(line string) (moves []int) {
	for _, c := range line {
		switch c { // Might be opposite.
		case '<':
			moves = append(moves, -1)
		case '>':
			moves = append(moves, 1)
		}
	}
	return
}

func printShaft(shaft shaftT, T int) {
	for i := 0; i < 20; i++ {
		line := ""
		for _, c := range fmt.Sprintf("%9b", shaft[(T-i+1000)%1000]) {
			switch c {
			case '1':
				line += "#"
			case '0':
				line += " "
			default:
				line += string(c)
			}
		}
		Log(line)
	}
}
