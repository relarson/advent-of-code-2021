package main

import (
	"fmt"
	"strconv"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

const (
	NOT_VISITABLE int = iota
	REPEAT
	VISITABLE
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day13/test_input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	return len(lines)
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day13/test_input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	return len(lines)
}
