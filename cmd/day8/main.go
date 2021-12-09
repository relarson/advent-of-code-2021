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
	lines, err := pkg.ReadLines("cmd/day8/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	targetNums := 0
	for _, line := range lines {
		_, outputs := parseInputLine(line)
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				targetNums++
			}
		}
	}

	return targetNums
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day8/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	return len(lines)
}

func parseInputLine(line string) ([]string, []string) {
	halves := strings.Split(line, " | ")
	signals := strings.Split(halves[0], " ")
	outputs := strings.Split(halves[1], " ")

	return signals, outputs
}
