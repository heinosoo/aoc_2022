package main

import (
	"fmt"
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

type tileT struct {
	next  [4]*tileT // >, v, <, ^
	wall  bool
	score int
}

func part1(lines Channel[string]) {
	tile, moves, _ := parseMap(lines)
	dir := 0
	for _, move := range moves {
		switch move {
		case "R":
			dir = (4 + dir + 1) % 4
		case "L":
			dir = (4 + dir - 1) % 4
		default:
			n, _ := strconv.Atoi(string(move))
			tile = tile.walk(n, dir)
		}
	}
	Log(tile.score + dir)
}

func (tile *tileT) walk(n, dir int) *tileT {
	for ; n > 0; n-- {
		next := tile.next[dir]
		if next.wall {
			break
		} else {
			tile = next
		}
	}
	return tile
}

func print(current *tileT, tiles [][]*tileT) {
	for _, row := range tiles {
		for _, tile := range row {
			if tile == nil {
				fmt.Print(" ")
			} else if tile == current {
				fmt.Print("o")
			} else if tile.wall {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		Log()
	}
}

func parseMap(lines Channel[string]) (start *tileT, moves []string, tiles [][]*tileT) {
	tiles2 := [][]*tileT{}
	y, maxX := 0, 0
	for line := range lines {
		maxX = Max(maxX, len(line)-1)
		row := []*tileT{}
		if line == "" {
			break
		}
		for x, c := range line {
			score := 1000*(y+1) + 4*(x+1)
			switch c {
			case ' ':
				row = append(row, nil)
			case '.':
				row = append(row, &tileT{wall: false, score: score})
			case '#':
				row = append(row, &tileT{wall: true, score: score})
			}
		}
		tiles2 = append(tiles2, row)
		y++
	}
	tiles = [][]*tileT{}
	for _, row := range tiles2 {
		for i := len(row); i <= maxX; i++ {
			row = append(row, nil)
		}
		tiles = append(tiles, row)
	}
	for y, row := range tiles {
		for x, tile := range row {
			if tile == nil {
				continue
			}
			for d := 1; ; d++ {
				if y-d < 0 {
					d -= len(tiles)
				}
				if tiles[y-d][x] != nil {
					tile.next[3] = tiles[y-d][x]
					break
				}
			}
			for d := 1; ; d++ {
				if y+d >= len(tiles) {
					d -= len(tiles)
				}
				if tiles[y+d][x] != nil {
					tile.next[1] = tiles[y+d][x]
					break
				}
			}
			for d := 1; ; d++ {
				if x-d < 0 {
					d -= len(row)
				}
				if tiles[y][x-d] != nil {
					tile.next[2] = tiles[y][x-d]
					break
				}
			}
			for d := 1; ; d++ {
				if x+d >= len(row) {
					d -= len(row)
				}
				if tiles[y][x+d] != nil {
					tile.next[0] = tiles[y][x+d]
					break
				}
			}
		}
	}
	for _, tile := range tiles[0] {
		if tile != nil {
			start = tile
			break
		}
	}
	n := ""
	for _, c := range <-lines {
		if c == 'L' || c == 'R' {
			moves = append(moves, n, string(c))
			n = ""
		} else {
			n += string(c)
		}
	}
	moves = append(moves, n)

	return
}
