package day11

import (
	. "adventofcode2022/tools"
	"fmt"
	"sort"
	"strings"
)

type monkeyOps struct {
	monkey                int
	startingItems         []int
	divisibleBy           int
	throwTrue, throwFalse int
	calcWorryLevel        func(worryLevel int) int
}

func getMonkeysOps() []monkeyOps {
	scanner := LoadFile("input/input11.txt")

	monkeys := []monkeyOps{}
	for scanner.Scan() {
		monkey := monkeyOps{}
		// get Mokey
		m := strings.Split(scanner.Text(), " ")
		monkey.monkey = Atoi(m[1][:len(m[1])-1])
		// Starting items
		scanner.Scan()
		monkey.startingItems = []int{}
		items := strings.Split(scanner.Text(), " ")
		for i := 4; i < len(items); i++ {
			nr := strings.Split(items[i], ",")
			monkey.startingItems = append(monkey.startingItems, Atoi(nr[0]))
		}
		// Operation
		scanner.Scan()
		ops := strings.Split(scanner.Text(), " ")
		if ops[7] == "old" {
			if ops[6] == "*" {
				monkey.calcWorryLevel = func(old int) int { return old * old }
			} else {
				monkey.calcWorryLevel = func(old int) int { return old + old }
			}
		} else {
			if ops[6] == "*" {
				monkey.calcWorryLevel = func(old int) int { return old * Atoi(ops[7]) }
			} else {
				monkey.calcWorryLevel = func(old int) int { return old + Atoi(ops[7]) }
			}
		}
		// Test
		scanner.Scan()
		monkey.divisibleBy = Atoi(strings.Split(scanner.Text(), " ")[5])
		// If true
		scanner.Scan()
		monkey.throwTrue = Atoi(strings.Split(scanner.Text(), " ")[9])
		// If false
		scanner.Scan()
		monkey.throwFalse = Atoi(strings.Split(scanner.Text(), " ")[9])

		monkeys = append(monkeys, monkey)
		scanner.Scan() // Empty line
	}
	return monkeys
}

func calcMonkeyBusiness(rounds int, monkeys []monkeyOps, reduceStress func(worryLevel int) int) int {

	inspections := make([]int, len(monkeys))
	for i := 1; i <= rounds; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.startingItems {
				worryLevel := reduceStress(monkey.calcWorryLevel(item))
				if worryLevel%monkey.divisibleBy == 0 {
					monkeys[monkey.throwTrue].startingItems = append(monkeys[monkey.throwTrue].startingItems, worryLevel)
				} else {
					monkeys[monkey.throwFalse].startingItems = append(monkeys[monkey.throwFalse].startingItems, worryLevel)
				}
				inspections[m]++
			}
			monkeys[m].startingItems = []int{} // Empties monkey list

		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	return inspections[0] * inspections[1]
}

// find Least Common Multiple
func lcm(monkeys []monkeyOps) int {
	result := 1
	for _, monkey := range monkeys {
		result *= monkey.divisibleBy
	}
	return result
}

func Ex1() {
	fmt.Println("Ex1: ", calcMonkeyBusiness(20, getMonkeysOps(), func(worryLevel int) int { return worryLevel / 3 }))
}
func Ex2() {
	monkeys := getMonkeysOps()
	fmt.Println("Ex2: ", calcMonkeyBusiness(10000, monkeys, func(worryLevel int) int { return worryLevel % lcm(monkeys) }))
}
