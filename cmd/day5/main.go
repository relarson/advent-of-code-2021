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
	lines, err := pkg.ReadLines("cmd/day5/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	counts := make(map[int]map[int]int)

	for _, line := range lines {
		pairs := strings.Split(line, " -> ")
		p := strings.Split(pairs[0], ",")
		x1, _ := strconv.Atoi(p[0])
		y1, _ := strconv.Atoi(p[1])

		p = strings.Split(pairs[1], ",")
		x2, _ := strconv.Atoi(p[0])
		y2, _ := strconv.Atoi(p[1])

		if x1 == x2 {
			// vert line
			if y2 < y1 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			//println("Adding vertical in column: ", x1, " starting at ", y1, " and ending at ", y2)
			for r := y1; r <= y2; r++ {
				addToMap(&counts, r, x1)
			}
			//printMap(counts)
		} else if y1 == y2 {
			// ys must be equal since we only have vert or horiz lines
			// horiz line
			if x2 < x1 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}
			//println("Adding horizontal in row: ", y1, " starting at ", x1, " and ending at ", x2)
			for c := x1; c <= x2; c++ {
				addToMap(&counts, y1, c)
			}
			//printMap(counts)
		} else {
			// ignore diagonals
			//println("Ignore diagonal " + line)
		}
	}

	return countHotSpots(counts)

}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day5/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}
	return len(lines)
}

func addToMap(counts *map[int]map[int]int, r int, c int) {
	row, ok := (*counts)[r]
	if ok {
		_, ok := row[c]
		if ok {
			(*counts)[r][c]++
		} else {
			(*counts)[r][c] = 1
		}
	} else {
		(*counts)[r] = make(map[int]int)
		(*counts)[r][c] = 1
	}
}

func countHotSpots(counts map[int]map[int]int) int {
	twoOrMore := 0
	for _, row := range counts {
		for _, val := range row {
			if val >= 2 {
				twoOrMore++
			}
		}
	}

	return twoOrMore
}

func printMap(counts map[int]map[int]int) {
	for r := 0; r < 10; r++ {
		row, ok := counts[r]
		if ok {
			for c := 0; c < 10; c++ {
				val, ok := row[c]
				if ok {
					print(val, " ")
				} else {
					print(". ")
				}
			}
		} else {
			print(". . . . . . . . . ")
		}
		println("")
	}
}
