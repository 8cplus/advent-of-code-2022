package day9

import (
	. "adventofcode2022/tools"
	"fmt"
)

/****************** Auxiliary Debugging functions ******************/

func findXYinSlice(xy XY, snake []XY) int {
	for i, snakeXY := range snake {
		if snakeXY == xy {
			return i
		}
	}
	return -1
}

func printSnake(snake []XY) {
	minX := snake[0].X
	minY := snake[0].Y
	maxX := snake[0].X
	maxY := snake[0].Y
	for _, xy := range snake {
		maxX = Max(xy.X, maxX)
		maxY = Max(xy.Y, maxY)
		minX = Min(xy.X, minX)
		minY = Min(xy.Y, minY)
	}

	for i := maxY + 1; i >= minY-1; i-- {
		for j := minX - 1; j <= maxX+1; j++ {
			if p := findXYinSlice(XY{X: j, Y: i}, snake); p >= 0 {
				if p == 0 {
					fmt.Print("H")
				} else {
					fmt.Print(p)
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(string("\033[37m"))
}

func printTravelGrid(travel map[XY]int) {
	var minX, maxX, minY, maxY int
	for xy := range travel {
		maxX = Max(xy.X, maxX)
		maxY = Max(xy.Y, maxY)
		minX = Min(xy.X, minX)
		minY = Min(xy.Y, minY)
	}

	for i := maxY; i >= minY; i-- {
		for j := minX; j <= maxX; j++ {
			if i == 0 && j == 0 {
				fmt.Print(string("\033[31m"), "X")
				continue
			}
			if travel[XY{X: j, Y: i}] > 0 {
				fmt.Print(string("\033[37m"), "#")
			} else {
				fmt.Print(string("\033[37m"), ".")
			}

		}
		fmt.Println()
	}
	fmt.Print(string("\033[37m"))
}
