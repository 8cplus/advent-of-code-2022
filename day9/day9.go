package day9

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

var moveHead = map[string]XY{"U": {X: 0, Y: 1}, "D": {X: 0, Y: -1}, "R": {X: 1, Y: 0}, "L": {X: -1, Y: 0}}

func headDirection(a, b XY) (c XY) {
	if a.X-b.X > 0 {
		c.X = 1
	} else if a.X-b.X < 0 {
		c.X = -1
	}
	if a.Y-b.Y > 0 {
		c.Y = 1
	} else if a.Y-b.Y < 0 {
		c.Y = -1
	}
	return c
}

func moveTail(h, t XY) XY {
	d := headDirection(h, t)
	if (h.X == t.X && h.AbsDist(t).Y > 1) || (h.Y == t.Y && h.AbsDist(t).X > 1) || h.AbsDist(t).X > 1 && h.AbsDist(t).Y > 1 {
		t.Add(d)
	} else if h.AbsDist(t).X > 1 {
		t.X += d.X
		t.Y = h.Y
	} else if h.AbsDist(t).Y > 1 {
		t.Y += d.Y
		t.X = h.X
	}
	return t
}

func moveSnake(snakeSize int) int {
	scanner := LoadFile("input/input9.txt")
	tailTravel := map[XY]int{{}: 1}
	snake := make([]XY, snakeSize)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		for i := 1; i <= Atoi(moves[1]); i++ {
			snake[0].Add(moveHead[moves[0]])
			for j := 0; j < len(snake)-1; j++ {
				snake[j+1] = moveTail(snake[j], snake[j+1])
			}
			tailTravel[snake[len(snake)-1]]++
		}
	}
	return len(tailTravel)
}

func Ex1andEx2() {
	fmt.Println("Ex1: ", moveSnake(2))
	fmt.Println("Ex2: ", moveSnake(10))
}
