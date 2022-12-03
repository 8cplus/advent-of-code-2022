package day3

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

// Converts a rune to an int based on 'a' = 1 and 'A' = 27
func char2Int(s rune) int {
	if s >= 'a' && s <= 'z' {
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
		for _, elem := range rucksack[:len(rucksack)/2] {
			if strings.ContainsRune(rucksack[len(rucksack)/2:], elem) {
				sum += char2Int(elem)
				break
			}
		}
	}
	fmt.Println("Ex1: ", sum)
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
	fmt.Println("Ex2: ", sum)
}
