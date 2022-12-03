package main

import (
	"bufio"
	"os"
	"strings"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	s := 0
	for badgePriority := nextBadge(scanner); badgePriority > 0; badgePriority = nextBadge(scanner) {
		s += badgePriority
	}
	return s
}

func nextBadge(scanner *bufio.Scanner) (badgePriority int) {
	if !scanner.Scan() {
		return
	}
	elfA := scanner.Text()
	scanner.Scan()
	elfB := scanner.Text()
	scanner.Scan()
	elfC := scanner.Text()

	for _, item := range elfA {
		if strings.ContainsRune(elfB, item) && strings.ContainsRune(elfC, item) {
			return priority(int(item))
		}
	}

	panic("No badge found.")

}
