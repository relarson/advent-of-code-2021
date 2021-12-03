package main

import (
	"fmt"
	"strconv"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day3/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	digits := len(lines[0])
	entries := len(lines)

	positionSums := make([]int, digits)

	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				positionSums[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0

	for i, sum := range positionSums {
		decimalValueToAdd := (1 << (digits - i - 1))
		if sum > (entries / 2) {
			// 1s were more common
			gamma += decimalValueToAdd
		} else {
			epsilon += decimalValueToAdd
		}
	}

	return gamma * epsilon
}

func problem2() int {
	return -1
}
