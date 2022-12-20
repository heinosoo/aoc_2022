package main

import (
	"strconv"

	. "github.com/heinosoo/aoc_2022"
)

type circle struct {
	val, len   int
	prev, next *circle
}

func (num *circle) get(N int) (ret *circle) {
	ret = num
	if N > 0 {
		for n := 0; n < N; n++ {
			ret = ret.next
		}
	} else if N < 0 {
		for n := 0; n < -N; n++ {
			ret = ret.prev
		}
	}
	return
}
func (num *circle) shift() {
	var target *circle
	m := num.val % (num.len - 1) // Stupid -1
	if m == 0 {
		return
	} else if m > 0 {
		target = num.get(m)
	} else {
		target = num.get(m - 1)
	}
	num.next.prev, num.prev.next = num.prev, num.next
	num.prev, num.next = target, target.next
	target.next, target.next.prev = num, num
}

func part1(lines Channel[string]) {
	numbers, zero := parseFile(lines, 1)
	for _, num := range numbers {
		num.shift()
	}
	Log(zero.get(1000).val + zero.get(2000).val + zero.get(3000).val)
}

func parseFile(lines Channel[string], key int) (numbers []*circle, zero *circle) {
	for line := range lines {
		val, _ := strconv.Atoi(line)
		numbers = append(numbers, &circle{val: val * key})
	}
	for i := range numbers {
		numbers[i].len = len(numbers)
		numbers[i].prev = numbers[(len(numbers)+i-1)%len(numbers)]
		numbers[i].next = numbers[(len(numbers)+i+1)%len(numbers)]
		if numbers[i].val == 0 {
			zero = numbers[i]
		}
	}
	return
}
