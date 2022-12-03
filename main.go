package main

import (
	"adventofcode2022/day3"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("################### Ex1 ###############################")
	day3.Ex1()
	fmt.Println("################### Ex2 ###############################")
	day3.Ex2()
	fmt.Println("\nTime: ", time.Since(start))

}
