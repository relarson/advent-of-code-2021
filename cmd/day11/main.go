package main

import (
	"fmt"
	"strconv"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

type Point struct {
	row    int
	column int
}

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	matrix, err := pkg.ReadIntMatrix("cmd/day11/input.txt", "")
	if err != nil {
		println(err.Error())
		return -1
	}

	flashes := 0
	for d := 0; d < 100; d++ {
		incrementAll(&matrix)
		flashes += handleFlashes(&matrix)
		resetFlashed(&matrix)
	}

	return flashes
}

func problem2() int {
	matrix, err := pkg.ReadIntMatrix("cmd/day11/input.txt", "")
	if err != nil {
		println(err.Error())
		return -1
	}

	loop := true
	day := 0
	targetFlashes := len(matrix) * len(matrix[0])

	for loop {
		day++
		incrementAll(&matrix)
		flashes := handleFlashes(&matrix)
		if flashes == targetFlashes {
			return day
		}
		resetFlashed(&matrix)
	}

	return -1
}

func incrementAll(matrix *[][]int) {
	for r := range *matrix {
		for c := range (*matrix)[r] {
			(*matrix)[r][c]++
		}
	}
}

func handleFlashes(matrix *[][]int) int {
	flashes := 0
	var toFlash []Point
	// build list of initial flashes
	for r := range *matrix {
		for c := range (*matrix)[r] {
			if (*matrix)[r][c] > 9 {
				toFlash = append(toFlash, Point{r, c})
			}
		}
	}

	for len(toFlash) > 0 {
		current := toFlash[0]
		toFlash = toFlash[1:]

		// ensure this didnt get flashed already between added to stack and when handled
		if (*matrix)[current.row][current.column] > 9 {
			// set sentinal so we dont add to it and can reset later
			(*matrix)[current.row][current.column] = -1
			flashes++
			readyToFlash := incrementNeighbors(current, matrix)
			toFlash = append(toFlash, readyToFlash...)
		}
	}

	return flashes
}

func incrementNeighbors(flashed Point, matrix *[][]int) []Point {
	toIncrement := []Point{
		{flashed.row - 1, flashed.column - 1},
		{flashed.row - 1, flashed.column},
		{flashed.row - 1, flashed.column + 1},
		{flashed.row, flashed.column - 1},
		{flashed.row, flashed.column + 1},
		{flashed.row + 1, flashed.column - 1},
		{flashed.row + 1, flashed.column},
		{flashed.row + 1, flashed.column + 1},
	}

	var newFlashers []Point
	for _, pt := range toIncrement {
		// check if this is a valid point
		if pt.row >= 0 && pt.row < len(*matrix) && pt.column >= 0 && pt.column < len((*matrix)[0]) {
			val := (*matrix)[pt.row][pt.column]
			if val != -1 && val < 10 {
				(*matrix)[pt.row][pt.column]++
			}
			// compare to 9 since it would have jsut become 10
			if val == 9 {
				newFlashers = append(newFlashers, pt)
			}
		}
	}
	return newFlashers
}

func resetFlashed(matrix *[][]int) {
	for r := range *matrix {
		for c := range (*matrix)[r] {
			if (*matrix)[r][c] == -1 {
				(*matrix)[r][c] = 0
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	println()
	for _, row := range matrix {
		for _, val := range row {
			print(val)
		}
		println()
	}
}
