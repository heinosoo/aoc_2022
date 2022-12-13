package main

import (
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

func part1(lines Channel[string]) {
	pairs := lines.Split(func(a string) bool { return a == "" })
	sum, i := 0, 0
	for pair := range pairs {
		i++
		left, right := NewChannelFromSlice([]rune(<-pair), 3), NewChannelFromSlice([]rune(<-pair), 3)
		lp := readPacket(left).packets[0]
		rp := readPacket(right).packets[0]
		r, _ := lp.compare(rp)
		if r {
			sum += i
		}
	}
	Log(sum)
}

type packet struct {
	value   int
	packets []packet
}

func readPacket(s Channel[rune]) packet {
	p := &packet{value: -1}
	numString := ""
	for c := range s {
		switch c {
		case '[':
			p.packets = append(p.packets, readPacket(s))
		case ']':
			num, e := strconv.Atoi(numString)
			numString = ""
			if e == nil {
				p.packets = append(p.packets, packet{value: num})
			}
			return *p
		case ',':
			num, e := strconv.Atoi(numString)
			numString = ""
			if e == nil {
				p.packets = append(p.packets, packet{value: num})
			}
		default:
			numString += string(c)
		}
	}
	return *p
}

func (left packet) compare(right packet) (bool, bool) {
	if left.value != -1 && right.value != -1 {
		if left.value != right.value {
			return left.value < right.value, true
		} else {
			return true, false
		}
	} else if left.value != -1 && right.value == -1 {
		return packet{value: -1, packets: []packet{{value: left.value}}}.compare(right)
	} else if left.value == -1 && right.value != -1 {
		return left.compare(packet{value: -1, packets: []packet{{value: right.value}}})
	} else {
		for i, rp := range right.packets {
			if i == len(left.packets) {
				return true, true
			} else {
				res, stop := left.packets[i].compare(rp)
				if stop {
					return res, stop
				}
			}
		}
		if len(left.packets) == len(right.packets) {
			return true, false
		}
		return false, true
	}
}
