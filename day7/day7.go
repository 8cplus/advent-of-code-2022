package day7

import (
	. "adventofcode2022/tools"
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

func applyCommands(commands []string, currentDir *directory) (remainingCommands []string, size int) {

	for len(commands) > 0 {
		cmd := strings.Split(commands[0], " ")
		if cmd[0][0:1] == "$" {
			if cmd[1] == "cd" {
				if cmd[2] == ".." { // cd ..
					return commands[1:], currentDir.sizeOfFiles
				} else { // cd folder
					subSize := 0
					commands, subSize = applyCommands(commands[1:], currentDir.subDir[cmd[2]])
					currentDir.sizeOfFiles += subSize
				}
			} else { // ls
				commands = commands[1:]
			}
		} else {
			if cmd[0] == "dir" { // directory
				currentDir.subDir[cmd[1]] = &directory{subDir: make(map[string]*directory)}
				commands = commands[1:]
			} else { // file
				currentDir.sizeOfFiles += Atoi(cmd[0])
				commands = commands[1:]
			}
		}
	}
	return nil, currentDir.sizeOfFiles
}

func Ex1andEx2() {
	scanner := LoadFile("input/input7.txt")
	scanner.Scan() // first cd /
	root := &directory{subDir: make(map[string]*directory)}
	commands := []string{}

	for scanner.Scan() {
		strIn := scanner.Text()
		commands = append(commands, strIn)
	}

	applyCommands(commands, root)

	fmt.Println("Ex1: ", root.sumSizesLessEqThan(100000))
	fmt.Println("Ex2: ", root.spaceToFree(70000000, 30000000-(70000000-root.sizeOfFiles)))
}
