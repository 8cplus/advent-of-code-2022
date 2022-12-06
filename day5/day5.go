package day5

import (
	. "adventofcode2022/tools"
	"fmt"
	"strconv"
	"strings"
)

// moves n crates from one pile to the other
func moveFuncEx1(move, from, to int, piles [][]byte) {
	for i := 0; i < move; i++ {
		piles[to] = append([]byte{piles[from][0]}, piles[to]...)
		piles[from] = piles[from][1:]
	}
}

// moves n crates from one pile to the other, mantaining the order of the moved crates
func moveFuncEx2(move, from, to int, piles [][]byte) {
	for i := move - 1; i >= 0; i-- {
		piles[to] = append([]byte{piles[from][i]}, piles[to]...)
	}
	piles[from] = piles[from][move:]
}

func crateMove(moveFunc func(int, int, int, [][]byte)) {
	scanner := LoadFile("input/input5.txt")
	var pilesOfCrates [][]byte = nil
	for scanner.Scan() {
		strIn := scanner.Text()

		if pilesOfCrates == nil {
			pilesOfCrates = make([][]byte, (len(strIn)+1)/4) // there's a crate every 4 positions
		}
		if _, err := strconv.Atoi(strIn[1:2]); err == nil {
			scanner.Scan() // gets the empty line after the crates and before the commands
			break
		}

		for i := 0; i < (len(strIn)+1)/4; i++ {
			if pilesOfCrates[i] == nil {
				pilesOfCrates[i] = []byte{}
			}
			if crate := strIn[i*4+1]; crate != ' ' {
				pilesOfCrates[i] = append(pilesOfCrates[i], crate)
			}
		}
	}

	for scanner.Scan() {
		strIn := scanner.Text()
		command := strings.Split(strIn, " ")
		// splited string - move 1 from 3 to 5
		moveFunc(Atoi(command[1]), Atoi(command[3])-1, Atoi(command[5])-1, pilesOfCrates)
	}

	for _, pileOfCrates := range pilesOfCrates {
		fmt.Printf("%c", pileOfCrates[0])
	}
}

func Ex1andEx2() {
	crateMove(moveFuncEx1)
	fmt.Println()
	crateMove(moveFuncEx2)
}
