package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.FormatInt(problem1(), 10))
	fmt.Println("Problem 2: " + strconv.FormatInt(problem2(), 10))
}

type Fish struct {
	startingTimer int
	lifespan      int
}

func problem1() int64 {
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

	return calculateFish(toProcess, 80)
}

func problem2() int64 {
	lines, err := pkg.ReadLines("cmd/day6/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	var toProcess []Fish

	fishStr := strings.Split(lines[0], ",")
	for _, str := range fishStr {
		val, _ := strconv.Atoi(str)
		// original fish all get 256d lifespan
		toProcess = append(toProcess, Fish{val, 256})
	}

	return calculateFish(toProcess, 256)
}

/* unused
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
*/

func calculateFish(toProcess []Fish, days int) int64 {
	popByTimer := make(map[int]int64, 9)

	for i := 0; i < 9; i++ {
		popByTimer[i] = 0
	}

	for _, fish := range toProcess {
		popByTimer[fish.startingTimer]++
	}

	for day := 1; day <= days; day++ {
		spawnAmt := popByTimer[0]
		for t := 0; t < 9; t++ {
			switch t {
			case 8:
				popByTimer[t] = spawnAmt // newly spawned
			case 6:
				popByTimer[t] = popByTimer[t+1] + spawnAmt // spawners reset
			default:
				popByTimer[t] = popByTimer[t+1]
			}
		}
	}

	var numberOfFish int64
	for i := 0; i < 9; i++ {
		numberOfFish += popByTimer[i]
	}

	return numberOfFish
}
