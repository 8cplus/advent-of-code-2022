package day2

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

const (
	Rock int = iota
	Paper
	Sissor
)

const (
	Lose int = iota
	Draw
	Win
)

func score(player int, otherPlayer int) int {
	winLoseTable := [][]int{{Rock, Sissor}, {Sissor, Paper}, {Paper, Rock}}

	for _, winLose := range winLoseTable {
		if player == winLose[0] && otherPlayer == winLose[1] {
			return player + 7
		}
	}
	if player == otherPlayer {
		return player + 4
	}
	return player + 1
}

func Ex1() {
	scanner := LoadFile("input/input2.txt")

	total := 0
	for scanner.Scan() {
		strIn := scanner.Text()
		line := strings.Split(strIn, " ")

		otherPlay := int(line[0][0]) - int('A')
		myPlay := int(line[1][0]) - int('X')
		total += score(myPlay, otherPlay)
	}
	fmt.Println("Ex1: ", total)
}

func Ex2() {
	scanner := LoadFile("input/input2.txt")

	total := 0
	for scanner.Scan() {
		strIn := scanner.Text()
		line := strings.Split(strIn, " ")

		myPlay := 0
		otherPlay := int(line[0][0]) - int('A')
		gameResult := int(line[1][0]) - int('X')
		switch gameResult {
		case Draw:
			myPlay = otherPlay
		case Win:
			myPlay = (otherPlay + 1) % 3
		case Lose:
			myPlay = (otherPlay + 2) % 3
		}
		total += score(myPlay, otherPlay)
	}
	fmt.Println("Ex2: ", total)
}
