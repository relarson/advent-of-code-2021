package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day10/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	totalScore := 0
	for _, line := range lines {
		completion, char := findCompletionOrCorruption(line)
		if completion.Len() == 0 {
			totalScore += scoreError(char)
		}
	}

	return totalScore
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day10/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	var scores []int
	for _, line := range lines {
		completion, _ := findCompletionOrCorruption(line)
		if completion.Len() != 0 {
			scores = append(scores, scoreCompletion(completion))
		}
	}
	sort.Ints(scores)

	mid := (len(scores) - 1) / 2

	return scores[mid]
}

func findCompletionOrCorruption(line string) (stack.Stack, string) {
	var currentlyOpen stack.Stack

	chars := strings.Split(line, "")

	for _, char := range chars {
		if strings.Contains("([{<", char) {
			currentlyOpen.Push(char)
		} else {
			shouldGetClosed := currentlyOpen.Pop().(string)
			switch shouldGetClosed {
			case "(":
				if char != ")" {
					return stack.Stack{}, char
				}
			case "[":
				if char != "]" {
					return stack.Stack{}, char
				}
			case "{":
				if char != "}" {
					return stack.Stack{}, char
				}
			case "<":
				if char != ">" {
					return stack.Stack{}, char
				}
			}
		}
	}
	return currentlyOpen, ""
}

func scoreError(invalid string) int {
	switch invalid {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func scoreCompletion(toClose stack.Stack) int {
	score := 0
	for toClose.Len() > 0 {
		char := toClose.Pop()
		switch char {
		case "(":
			score = score*5 + 1
		case "[":
			score = score*5 + 2
		case "{":
			score = score*5 + 3
		case "<":
			score = score*5 + 4
		}
	}
	return score
}
