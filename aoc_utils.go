package aoc_utils

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"golang.org/x/exp/constraints"
)

func GetFilename() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		return "input"
	}
}

func ReadFile(filename string, buffer chan string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer <- scanner.Text()
	}
	close(buffer)

	return
}

func Log(message ...any) {
	fmt.Println(message...)
}

func WaitForInput() {
	var wait string
	fmt.Scanln(&wait)
}

func WithTiming(prefix string, doSomething func()) {
	t := time.Now()
	doSomething()
	Log(prefix, time.Since(t))
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

func OperateOnMatrix[T any](matrix *[][]T, function func(T)) {
	M, N := len(*matrix), len((*matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			function((*matrix)[m][n])
		}
	}
}

func TransformMatrix[T any](matrix *[][]T, function func(T) T) {
	M, N := len(*matrix), len((*matrix)[0])
	for m := 0; m < M; m++ {
		for n := 0; n < N; n++ {
			(*matrix)[m][n] = function((*matrix)[m][n])
		}
	}
}

func InitializeMatrix[T any](value T, M, N int) (matrix [][]T) {
	for m := 0; m < M; m++ {
		matrix = append(matrix, make([]T, N))
		for n := 0; n < N; n++ {
			matrix[m][n] = value
		}
	}
	return
}
