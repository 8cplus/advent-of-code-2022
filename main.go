package main

import (
	"adventofcode2022/day3"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	/*fmt.Println("################### Day1 Ex1 ###############################")
	day1.Ex1()
	fmt.Println("################### Day1 Ex2 ###############################")
	day1.Ex2()
	fmt.Println("################### Day2 Ex1 ###############################")
	day2.Ex1()
	fmt.Println("################### Day2 Ex2 ###############################")
	day2.Ex2()*/
	fmt.Println("################### Day3 Ex1 ###############################")
	day3.Ex1()
	fmt.Println("################### Day3 Ex2 ###############################")
	day3.Ex2()
	/*fmt.Println("################### Day4 Ex1 ###############################")
	day4.Ex1()
	fmt.Println("################### Day4 Ex2 ###############################")
	day4.Ex2()*/

	fmt.Println("\nTime: ", time.Since(start))
}
