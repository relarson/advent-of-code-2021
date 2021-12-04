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
	lines, err := pkg.ReadLines("cmd/day3/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}
	oxgen := lines
	scrub := lines
	digits := len(lines[0])
	for i := 0; i < digits; i++ {
		if len(oxgen) > 1 {
			oxgen = filter(oxgen, i, true)
		}
		if len(scrub) > 1 {
			scrub = filter(scrub, i, false)
		}
	}
	// yolo on errors trust the input
	oxygen, _ := strconv.ParseInt(oxgen[0], 2, 64)
	scrubber, _ := strconv.ParseInt(scrub[0], 2, 64)
	return int(oxygen) * int(scrubber)
}

func filter(values []string, index int, filterToMostCommon bool) []string {
	var leadOnes []string
	var leadZeros []string

	for _, value := range values {
		if value[index] == '1' {
			leadOnes = append(leadOnes, value)
		} else {
			leadZeros = append(leadZeros, value)
		}
	}

	if filterToMostCommon {
		if len(leadOnes) >= len(leadZeros) {
			return leadOnes
		} else {
			return leadZeros
		}
	} else {
		if len(leadZeros) <= len(leadOnes) {
			return leadZeros
		} else {
			return leadOnes
		}
	}
}
