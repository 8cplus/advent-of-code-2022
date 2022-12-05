package main

import (
	"adventofcode2022/day5"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*fmt.Println("################### Day1 ###############################")
	day1.Ex1()
	day1.Ex2()
	fmt.Println("################### Day2 ###############################")
	day2.Ex1()
	day2.Ex2()
	fmt.Println("################### Day3 ###############################")
	day3.Ex1()
	day3.Ex2()
	fmt.Println("################### Day4 #########################")
	day4.Ex1andEx2()*/
	fmt.Println("################### Day5 #########################")
	day5.Ex1()
	day5.Ex2()

	fmt.Println("\nTime: ", time.Since(start))
}
