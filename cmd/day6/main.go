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

type Fish struct {
	startingTimer int
	lifespan      int
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day6/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	var toProcess []Fish

	fishStr := strings.Split(lines[0], ",")
	for _, str := range fishStr {
		val, _ := strconv.Atoi(str)
		// original fish all get 80d lifespan
		toProcess = append(toProcess, Fish{val, 80})
	}

	return processFish(toProcess)
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day6/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	return len(lines)
}

func processFish(toProcess []Fish) int {
	numberOfFish := len(toProcess)

	for len(toProcess) > 0 {
		// pop from queue
		currentFish := toProcess[0]
		toProcess = toProcess[1:]

		timer := currentFish.startingTimer
		for currentFish.lifespan > 0 {
			timeToSpawn := timer + 1
			currentFish.lifespan -= timeToSpawn
			if currentFish.lifespan >= 0 {
				// fish lives to spawn
				timer = 6 // reset timer
				// spawn fish with a timer of 8 that has the remaining lifespan
				toProcess = append(toProcess, Fish{8, currentFish.lifespan})
				numberOfFish++
			}
		}
	}

	return numberOfFish
}
