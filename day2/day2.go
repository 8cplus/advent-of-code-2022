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

func wins(r1 int, r2 int) int {

	score := r1 + 1

	if r1 == Rock && r2 == Sissor {
		return score + 6
	} else if r1 == Sissor && r2 == Paper {
		return score + 6
	} else if r1 == Paper && r2 == Rock {
		return score + 6
	} else if r1 == r2 {
		return score + 3
	}
	return score
}

func Ex1() {
	scanner := LoadFile("input/input2.txt")

	total := 0
	for scanner.Scan() {
		strIn := scanner.Text()
		line := strings.Split(strIn, " ")

		otherPlay := int(line[0][0]) - int('A')
		myPlay := int(line[1][0]) - int('X')
		total += wins(myPlay, otherPlay)
	}
	fmt.Println("Ex1: ", total)
}

func Ex2() {
	scanner := LoadFile("input/input2.txt")

	total := 0
	for scanner.Scan() {
		strIn := scanner.Text()
		line := strings.Split(strIn, " ")

		play := 0
		otherPlay := int(line[0][0]) - int('A')
		myPlay := int(line[1][0]) - int('X')
		if myPlay == Draw {
			play = otherPlay
		} else if myPlay == Win {
			play = (otherPlay + 1) % 3
		} else if myPlay == Lose {
			play = (otherPlay + 2) % 3
		}

		total += wins(play, otherPlay)
	}
	fmt.Println("Ex2: ", total)
}
