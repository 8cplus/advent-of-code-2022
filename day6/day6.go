package day6

import (
	. "adventofcode2022/tools"
	"fmt"
)

func ex1AndEx2(packetsize int) int {
	scanner := LoadFile("input/input6.txt")
	scanner.Scan()
	strIn := scanner.Text()

	for i := packetsize; i < len(strIn); i++ {

		found := true
	out:
		for j := i - packetsize; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if strIn[j] == strIn[k] {
					found = false
					break out
				}
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func Ex1andEx2() {
	fmt.Println("Ex1: ", ex1AndEx2(4))
	fmt.Println("Ex2: ", ex1AndEx2(14))
}
