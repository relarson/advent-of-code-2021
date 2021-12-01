package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLinesOfInts("cmd/day1/problem1_input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	previous := math.MaxInt
	increases := 0
	for _, depth := range lines {
		if depth > previous {
			increases++
		}
		previous = depth
	}

	return increases
}

func problem2() int {
	// uses same input
	lines, err := pkg.ReadLinesOfInts("cmd/day1/problem1_input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	previous := math.MaxInt
	increases := 0
	for i := 0; i < len(lines)-2; i++ {
		threeSum := lines[i] + lines[i+1] + lines[i+2]
		if threeSum > previous {
			increases++
		}
		previous = threeSum
	}

	return increases
}
