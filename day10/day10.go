package day10

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

type cpu struct {
	cycle int
	x     int
}

func checkCpuState(cycle int, cpuState []cpu) int {
	for i := 0; i < len(cpuState); i++ {
		if cycle == cpuState[i].cycle {
			return cpuState[i].x
		}
		if cycle < cpuState[i].cycle {
			return cpuState[i-1].x
		}
	}
	return cpuState[len(cpuState)-1].x
}

func drawSprite(x, sprite int) bool {
	if sprite-1 == x || sprite == x || sprite+1 == x {
		return true
	}
	return false
}

func Ex1andEx2() {
	scanner := LoadFile("input/input10.txt")

	cycle := 1
	x := 1
	cpuState := []cpu{{x: 1, cycle: 1}}

	for scanner.Scan() {
		strIn := scanner.Text()
		if strIn == "noop" {
			cycle++
			continue
		}
		x += Atoi(strings.Split(strIn, " ")[1])
		cycle += 2
		cpuState = append(cpuState, cpu{cycle: cycle, x: x})
	}
	signalStrenth := 0
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		signalStrenth += c * checkCpuState(c, cpuState)
	}
	fmt.Println("Ex1: ", signalStrenth)
	// Ex2
	for pixel := 0; pixel < 240; pixel++ {
		if pixel > 0 && pixel%40 == 0 {
			fmt.Println()
		}
		if drawSprite(pixel%40, checkCpuState(pixel+1, cpuState)) {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
}
