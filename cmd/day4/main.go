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
	lines, err := pkg.ReadLines("cmd/day4/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	var called []int
	for _, num := range strings.Split(lines[0], ",") {
		val, _ := strconv.Atoi(num)
		called = append(called, val)
	}

	fastestM := len(called) + 1
	scoreOfFastest := -1
	for i := 2; i < len(lines); i += 6 {
		var board Board
		for r := 0; r < 5; r++ {
			for _, num := range strings.Fields(lines[i+r]) {
				val, _ := strconv.Atoi(num)
				board.values = append(board.values, val)
			}
		}
		moves, score := board.Play(called)
		if moves < fastestM {
			fastestM = moves
			scoreOfFastest = score
		}
	}
	return scoreOfFastest
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day4/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	var called []int
	for _, num := range strings.Split(lines[0], ",") {
		val, _ := strconv.Atoi(num)
		called = append(called, val)
	}

	slowestM := 0
	scoreOfSlowest := -1
	for i := 2; i < len(lines); i += 6 {
		var board Board
		for r := 0; r < 5; r++ {
			for _, num := range strings.Fields(lines[i+r]) {
				val, _ := strconv.Atoi(num)
				board.values = append(board.values, val)
			}
		}
		moves, score := board.Play(called)
		if moves > slowestM {
			slowestM = moves
			scoreOfSlowest = score
		}
	}
	return scoreOfSlowest
}
