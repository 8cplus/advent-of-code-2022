package day11

import (
	. "adventofcode2022/tools"
	"fmt"
)

func Ex1andEx2() {
	scanner := LoadFile("input/input11.txt")

	for scanner.Scan() {
		strIn := scanner.Text()
		fmt.Print(strIn)
	}
}
