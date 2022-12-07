package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	root := emptyFolder(nil)
	current := root

	for scanner.Scan(); scanner.Text() != ""; {
		current = parseNextCmd(scanner, root, current)
	}

	_, smallSum := root.smallFolders(100000)
	return smallSum
}

type folder struct {
	parent  *folder
	folders map[string]*folder
	files   map[string]int
}

func emptyFolder(parent *folder) *folder {
	return &folder{parent: parent, folders: make(map[string]*folder), files: make(map[string]int)}
}
func (f *folder) addFolder(name string) {
	f.folders[name] = emptyFolder(f)
}
func (f *folder) addFile(name, size string) {
	s, _ := strconv.Atoi(size)
	f.files[name] = s
}
func (f folder) smallFolders(limit int) (size, smallSum int) {
	for _, fileSize := range f.files {
		size += fileSize
	}
	for _, subfolder := range f.folders {
		subSize, subSmallSum := subfolder.smallFolders(limit)
		size += subSize
		smallSum += subSmallSum
	}
	if size <= limit {
		smallSum += size
	}
	return
}

func parseNextCmd(scanner *bufio.Scanner, root, current *folder) *folder {
	cmd := strings.Split(scanner.Text(), " ")
	if cmd[0] != "$" {
		panic("Not a command")
	}
	switch cmd[1] {
	case "ls":
		func() {
			for scanner.Scan() {
				cmd := strings.Split(scanner.Text(), " ")
				switch cmd[0] {
				case "$":
					return
				case "dir":
					current.addFolder(cmd[1])
				default:
					current.addFile(cmd[1], cmd[0])
				}
			}
		}()
	case "cd":
		dir := cmd[2]
		scanner.Scan()
		switch dir {
		case "/":
			return root
		case "..":
			return current.parent
		default:
			return current.folders[dir]
		}
	}
	return current
}
