package main

import (
	"fmt"
	"sort"
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
	matrix, err := pkg.ReadIntMatrix("cmd/day9/input.txt", "")
	if err != nil {
		println(err.Error())
		return -1
	}

	lowPoints := findLowPoints(matrix)

	var sizes []int
	for _, point := range lowPoints {
		basinSize := findBasinSize(point, matrix)
		sizes = append(sizes, basinSize)
	}
	sort.Ints(sizes)

	lastIndex := len(sizes) - 1

	return sizes[lastIndex] * sizes[lastIndex-1] * sizes[lastIndex-2]
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

func findBasinSize(basinOrigin Point, matrix [][]int) int {
	settled := make(map[Point]bool, 0)
	var queue []Point

	queue = append(queue, basinOrigin)

	for len(queue) > 0 {
		cp := queue[0]
		queue = queue[1:]
		cv := matrix[cp.row][cp.column]
		neighbors := buildNeighbors(cp, matrix)
		for _, n := range neighbors {
			_, ok := settled[n]
			if ok {
				continue
			}
			nv := matrix[n.row][n.column]
			if nv > cv && nv < 9 {
				queue = append(queue, n)
			}
		}
		settled[cp] = true
	}

	return len(settled)
}

func buildNeighbors(pt Point, matrix [][]int) []Point {
	var neighbors []Point

	r := pt.row
	c := pt.column

	if r > 0 {
		neighbors = append(neighbors, Point{r - 1, c})
	}

	if r < (len(matrix) - 1) {
		neighbors = append(neighbors, Point{r + 1, c})
	}

	if c > 0 {
		neighbors = append(neighbors, Point{r, c - 1})
	}

	if c < (len(matrix[0]) - 1) {
		neighbors = append(neighbors, Point{r, c + 1})
	}

	return neighbors
}
