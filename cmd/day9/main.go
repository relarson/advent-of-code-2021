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
	matrix, err := pkg.ReadIntMatrix("cmd/day9/input.txt", "")
	if err != nil {
		println(err.Error())
		return -1
	}

	lowPoints := findLowPoints(matrix)

	totalRisk := 0
	for _, point := range lowPoints {
		val := matrix[point.row][point.column]
		totalRisk += val + 1
	}

	return totalRisk
}

func problem2() int {
	matrix, err := pkg.ReadIntMatrix("cmd/day9/test_input.txt", "")
	if err != nil {
		println(err.Error())
		return -1
	}

	return len(matrix)
}

func findLowPoints(matrix [][]int) []Point {
	var lowPoints []Point
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			current := matrix[r][c]
			if r > 0 {
				if current >= matrix[r-1][c] {
					continue
				}
			}

			if r < (len(matrix) - 1) {
				if current >= matrix[r+1][c] {
					continue
				}
			}

			if c > 0 {
				if current >= matrix[r][c-1] {
					continue
				}
			}

			if c < (len(matrix[0]) - 1) {
				if current >= matrix[r][c+1] {
					continue
				}
			}
			lowPoints = append(lowPoints, Point{r, c})
		}
	}
	return lowPoints
}
