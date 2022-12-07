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

func (d *directory) applyCommands(scanner *bufio.Scanner) (size int) {
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if cmd[0][0:1] == "$" {
			if cmd[1] == "cd" && cmd[2] == ".." { // cd ..
				return d.sizeOfFiles
			} else if cmd[1] == "cd" { // cd folder
				d.sizeOfFiles += d.subDir[cmd[2]].applyCommands(scanner)
			}
		} else if cmd[0] == "dir" { // directory
			d.subDir[cmd[1]] = &directory{subDir: make(map[string]*directory)}
		} else { // file
			d.sizeOfFiles += Atoi(cmd[0])
		}
	}
	return d.sizeOfFiles
}

func Ex1andEx2() {
	scanner := LoadFile("input/input7.txt")
	scanner.Scan() // cd /
	root := &directory{subDir: make(map[string]*directory)}

	root.applyCommands(scanner)
	fmt.Println("Ex1: ", root.sumSizesLessEqThan(100000))
	fmt.Println("Ex2: ", root.spaceToFree(70000000, 30000000-(70000000-root.sizeOfFiles)))
}
