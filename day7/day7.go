package day7

import (
	. "adventofcode2022/tools"
	"bufio"
	"fmt"
	"strings"
)

type directory struct {
	subDir      map[string]*directory
	sizeOfFiles int
}

func (d *directory) sumSizesLessEqThan(less int) int {
	size := 0
	for _, sub := range d.subDir {
		size += sub.sumSizesLessEqThan(less)
	}
	if d.sizeOfFiles <= less {
		return size + d.sizeOfFiles
	}
	return size
}

func (d *directory) spaceToFree(smallestDir, spaceNeeded int) int {
	for _, sub := range d.subDir {
		smallestDir = sub.spaceToFree(smallestDir, spaceNeeded)
	}
	if d.sizeOfFiles <= smallestDir && d.sizeOfFiles >= spaceNeeded {
		return d.sizeOfFiles
	}
	return smallestDir
}

func applyCommands(scanner *bufio.Scanner, currentDir *directory) (size int) {
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if cmd[0][0:1] == "$" {
			if cmd[1] == "cd" {
				if cmd[2] == ".." { // cd ..
					return currentDir.sizeOfFiles
				} else { // cd folder
					currentDir.sizeOfFiles += applyCommands(scanner, currentDir.subDir[cmd[2]])
				}
			}
		} else {
			if cmd[0] == "dir" { // directory
				currentDir.subDir[cmd[1]] = &directory{subDir: make(map[string]*directory)}
			} else { // file
				currentDir.sizeOfFiles += Atoi(cmd[0])
			}
		}
	}
	return currentDir.sizeOfFiles
}

func Ex1andEx2() {
	scanner := LoadFile("input/input7.txt")
	scanner.Scan() // first cd /
	root := &directory{subDir: make(map[string]*directory)}

	applyCommands(scanner, root)
	fmt.Println("Ex1: ", root.sumSizesLessEqThan(100000))
	fmt.Println("Ex2: ", root.spaceToFree(70000000, 30000000-(70000000-root.sizeOfFiles)))
}
