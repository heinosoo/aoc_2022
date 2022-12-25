package main

import (
	"fmt"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	s := 0
	for line := range lines {
		s += desnafu(line)
	}
	Log(snafu(s))
}

func snafu(n int) (s string) {
	reversed := ""
	for co := 0; n != 0 || co != 0; {
		co += n % 5
		n /= 5
		switch co % 5 {
		case 3:
			co += 6
			reversed += "="
		case 4:
			co += 5
			reversed += "-"
		default:
			reversed += fmt.Sprint(co % 5)
		}
		co /= 5
	}
	for i := len(reversed) - 1; i >= 0; i-- {
		s += string(reversed[i])
	}
	return
}

func desnafu(s string) (n int) {
	p := 1
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '0':
		case '1':
			n += p
		case '2':
			n += 2 * p
		case '-':
			n -= p
		case '=':
			n -= 2 * p
		}
		p *= 5
	}
	return
}
