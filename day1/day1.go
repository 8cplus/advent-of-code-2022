package day1

import (
	. "adventofcode2022/tools"
	"fmt"
	"sort"
)

func load() [][]int {

	scanner := LoadFile("input/input1.txt")

	var listOfelfs [][]int
	listOfelfs = append(listOfelfs, []int{})

	for scanner.Scan() {
		strIn := scanner.Text()
		if strIn == "" {
			listOfelfs = append(listOfelfs, []int{})
		} else {
			cal := Atoi(strIn)
			lastIndex := len(listOfelfs) - 1
			listOfelfs[lastIndex] = append(listOfelfs[lastIndex], cal)
		}
	}

	return listOfelfs
}

func Ex1() {

	listOfelfs := load()

	max := 0
	for _, elf := range listOfelfs {
		sum := 0
		for _, calories := range elf {
			sum += calories
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println("Ex1: ", max)
}

func Ex2() {
	listOfElfs := load()

	sumOfeachElf := make([]int, len(listOfElfs))
	for i := range listOfElfs {
		sum := 0
		for _, cal := range listOfElfs[i] {
			sum += cal
		}
		sumOfeachElf[i] = sum
	}
	sort.Ints(sumOfeachElf)
	size := len(sumOfeachElf)

	fmt.Println("Ex2: ", sumOfeachElf[size-3]+sumOfeachElf[size-2]+sumOfeachElf[size-1])
}
