package day4

import (
	. "adventofcode2022/tools"
	"fmt"
	"strings"
)

type section struct {
	start int
	end   int
}

func overlapAll(s1 section, s2 section) bool {
	if s1.start <= s2.start && s1.end >= s2.end {
		return true
	}
	return false
}

func Ex1andEx2() {
	scanner := LoadFile("input/input4.txt")

	var sumOverlapAll, sumOverlapAtLeastOne int
	for scanner.Scan() {
		strIn := scanner.Text()
		pair := strings.Split(strIn, ",")
		s1 := strings.Split(pair[0], "-")
		section1 := section{Atoi(s1[0]), Atoi(s1[1])}
		s2 := strings.Split(pair[1], "-")
		section2 := section{Atoi(s2[0]), Atoi(s2[1])}

		if overlapAll(section1, section2) || overlapAll(section2, section1) {
			sumOverlapAll++
		}
		if section1.start <= section2.end && section1.end >= section2.start {
			sumOverlapAtLeastOne++
		}
	}
	fmt.Println("Ex1: ", sumOverlapAll)
	fmt.Println("Ex2: ", sumOverlapAtLeastOne)
}
