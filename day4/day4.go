package day4

import (
	. "adventofcode2022/tools"
	"fmt"
)

func Ex1() {
	scanner := LoadFile("input/input4.txt")

	for scanner.Scan() {
		strIn := scanner.Text()
		fmt.Println(strIn)
	}
}

func Ex2() {
	scanner := LoadFile("input/input4.txt")

	for scanner.Scan() {
		strIn := scanner.Text()
		fmt.Println(strIn)
	}

}
