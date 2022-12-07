package main

import (
	"bufio"
	"os"

	aoc_utils "github.com/heinosoo/aoc_2022"
)

func part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := emptyFolder(nil)
	current := root

	for scanner.Scan(); scanner.Text() != ""; {
		current = parseNextCmd(scanner, root, current)
	}

	size, _ := root.smallFolders(0)
	_, smallest := root.findSmallest(size - 40000000)

	return smallest
}

func (f folder) findSmallest(limit int) (size, smallest int) {
	smallest = 2378549230482948375
	for _, fileSize := range f.files {
		size += fileSize
	}
	for _, subfolder := range f.folders {
		subSize, subSmallest := subfolder.findSmallest(limit)
		size += subSize
		if subSmallest > limit {
			smallest = aoc_utils.Min(smallest, subSmallest)
		}
	}
	smallest = aoc_utils.Min(smallest, size)
	return
}
