package day3

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

// Convertes a rune to an int based on 'a' = 1 and 'A' = 27
func char2Int(s rune) int {
	if s >= 97 && s <= 122 {
		return int(s - 96)
	} else {
		return int(s - 65 + 27)
	}
}

func Ex1() {
	scanner := LoadFile("input/input3.txt")

	sum := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		compart1 := rucksack[:len(rucksack)/2]
		compart2 := rucksack[len(rucksack)/2:]
		for _, elem := range compart1 {
			if strings.ContainsRune(compart2, elem) {
				repeated := char2Int(elem)
				sum += repeated
				break
			}
		}
	}
	fmt.Println("Result1: ", sum)
}

func Ex2() {
	scanner := LoadFile("input/input3.txt")

	var i, sum int
	var rucksack [3]string
	for scanner.Scan() {
		rucksack[i%3] = scanner.Text()
		if i%3 == 2 {
			for _, elem := range rucksack[0] {
				if strings.ContainsRune(rucksack[1], elem) && strings.ContainsRune(rucksack[2], elem) {
					sum += char2Int(elem)
					break
				}
			}
		}
		i++
	}
	fmt.Println("Result2: ", sum)
}
