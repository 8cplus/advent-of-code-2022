package day6

import (
	. "adventofcode2022/tools"
	"fmt"
)

func findMarker(strIn string, packetSize, pos int) int {
	for i := 0; i < packetSize; i++ {
		for j := i + 1; j < packetSize; j++ {
			if strIn[i] == strIn[j] {
				return findMarker(strIn[1:], packetSize, pos+1)
			}
		}
	}
	return pos
}

func Ex1andEx2() {
	scanner := LoadFile("input/input6.txt")
	scanner.Scan()

	fmt.Println("Ex1: ", findMarker(scanner.Text(), 4, 4))
	fmt.Println("Ex2: ", findMarker(scanner.Text(), 14, 14))
}
