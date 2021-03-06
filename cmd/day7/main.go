package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	crabs, err := pkg.ReadCommaDelimInts("cmd/day7/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	return calcFuel(crabs, median(crabs))
}

func problem2() int {
	crabs, err := pkg.ReadCommaDelimInts("cmd/day7/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	floored, ceiled := mean(crabs)

	floorFuel := calcFuelProgressive(crabs, floored)
	ceilFuel := calcFuelProgressive(crabs, ceiled)

	if floorFuel <= ceilFuel {
		return floorFuel
	} else {
		return ceilFuel
	}
}

func median(crabs []int) int {
	sort.Ints(crabs)
	count := len(crabs)

	if count%2 == 1 {
		return crabs[count/2]
	} else {
		return (crabs[(count-1)/2] + crabs[(count+1)/2]) / 2
	}
}

func mean(crabs []int) (int, int) {
	sum := 0
	for _, crab := range crabs {
		sum += crab
	}
	mean := float64(sum) / float64(len(crabs))
	return int(math.Floor(mean)), int(math.Ceil(mean))
}

func calcFuel(crabs []int, target int) int {
	total := 0
	for _, crab := range crabs {
		if crab > target {
			total += crab - target
		} else {
			total += target - crab
		}
	}
	return total
}

func calcFuelProgressive(crabs []int, target int) int {
	println(target)
	total := 0
	for _, crab := range crabs {
		if crab > target {
			diff := crab - target
			total += (diff * (diff + 1)) / 2
		} else {
			diff := target - crab
			total += (diff * (diff + 1)) / 2
		}
	}
	return total
}
