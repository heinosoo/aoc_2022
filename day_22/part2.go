package main

import (
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

func part2(lines Channel[string]) {
	tile, moves, tiles := parseMap(lines)
	wrap(tiles)
	dir := 0
	for _, move := range moves {
		switch move {
		case "R":
			dir = (4 + dir + 1) % 4
		case "L":
			dir = (4 + dir - 1) % 4
		default:
			n, _ := strconv.Atoi(string(move))
			tile, dir = tile.step(n, dir)
		}
	}
	Log(tile.score + dir)
}

func (tile *tileT) step(n, dir int) (*tileT, int) {
	for ; n > 0; n-- {
		next := tile.next[dir]
		if next.wall {
			break
		} else {
			for d2, prev := range next.next {
				if prev == tile {
					dir = (d2 + 2) % 4
				}
			}
			tile = next
		}
	}
	return tile, dir
}

func wrap(tiles [][]*tileT) {
	N := len(tiles) / 4
	Log(N)
	for i := 0; i < N; i++ {
		// example...
		// a1, b1 := tiles[0][2*N+i], tiles[N][N-1-i]
		// a2, b2 := tiles[N][N+i], tiles[i][2*N]
		// a3, b3 := tiles[2*N-1][i], tiles[3*N-1][3*N-1-i]
		// a4, b4 := tiles[2*N-1][N+i], tiles[3*N-1-i][2*N]
		// a5, b5 := tiles[2*N-1-i][3*N-1], tiles[2*N][3*N+i]
		// a6, b6 := tiles[i][3*N-1], tiles[3*N-1-i][4*N-1]
		// a7, b7 := tiles[3*N-1][3*N+i], tiles[2*N-1-i][0]
		// a1.next[3], b1.next[3] = b1, a1
		// a2.next[3], b2.next[2] = b2, a2
		// a3.next[1], b3.next[1] = b3, a3
		// a4.next[1], b4.next[2] = b4, a4
		// a5.next[0], b5.next[3] = b5, a5
		// a6.next[0], b6.next[0] = b6, a6
		// a7.next[1], b7.next[2] = b7, a7

		a34, a47, a25 := tiles[0][N+i], tiles[0][2*N+i], tiles[2*N][i]
		a16, a78, b47 := tiles[N-1][2*N+i], tiles[3*N-1][2*N-1-i], tiles[4*N-1][i]
		a23, b25, b23, b34 := tiles[N-1-i][N], tiles[N+i][N], tiles[2*N+i][0], tiles[3*N+i][0]
		a67, b16, b67, b78 := tiles[N-1-i][3*N-1], tiles[N+i][2*N-1], tiles[2*N+i][2*N-1], tiles[4*N-1-i][N-1]

		a34.next[3], b34.next[2] = b34, a34
		a47.next[3], b47.next[1] = b47, a47
		a25.next[3], b25.next[2] = b25, a25
		a16.next[1], b16.next[0] = b16, a16
		a78.next[1], b78.next[0] = b78, a78
		a23.next[2], b23.next[2] = b23, a23
		a67.next[0], b67.next[0] = b67, a67
	}
}
