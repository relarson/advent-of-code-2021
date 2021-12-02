package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day2/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	depth := 0
	horiz := 0

	for _, line := range lines {
		components := strings.Split(line, " ")
		direction := components[0]
		distance, err := strconv.Atoi(components[1])
		if err != nil {
			return -1
		}

		switch direction {
		case "up":
			distance *= -1
			fallthrough
		case "down":
			depth += distance
		case "forward":
			horiz += distance
		default:
			return -1
		}
	}

	return depth * horiz
}

func problem2() int {
	return -1
}
