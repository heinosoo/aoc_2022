package aoc_utils

import (
	"os"

	"golang.org/x/exp/constraints"
)

func GetFilename() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return "input"
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MinOf[T constraints.Ordered](vars ...T) (T, int) {
	min := vars[0]
	i := 0

	for j, n := range vars {
		if min > n {
			min, i = n, j
		}
	}

	return min, i
}

func MaxOf[T constraints.Ordered](vars ...T) (T, int) {
	max := vars[0]
	i := 0

	for j, n := range vars {
		if max < n {
			max, i = n, j
		}
	}

	return max, i
}

func Sum[T constraints.Integer | constraints.Float](vars ...T) T {
	var s T
	for _, n := range vars {
		s += n
	}
	return s
}