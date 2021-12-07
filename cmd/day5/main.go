package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

type Point struct {
	x int
	y int
}

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
		p, q := parsePoints(line)

		if p.x == q.x || p.y == q.y {
			points := pointsOnLine(p, q)
			for _, pt := range points {
				addToMap(&counts, pt)
			}
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

	counts := make(map[int]map[int]int)

	for _, line := range lines {
		p, q := parsePoints(line)

		points := pointsOnLine(p, q)
		for _, pt := range points {
			addToMap(&counts, pt)
		}
	}

	return countHotSpots(counts)
}

func parsePoints(line string) (Point, Point) {
	pairs := strings.Split(line, " -> ")
	pStr := strings.Split(pairs[0], ",")
	x1, _ := strconv.Atoi(pStr[0])
	y1, _ := strconv.Atoi(pStr[1])
	p := Point{x1, y1}

	qStr := strings.Split(pairs[1], ",")
	x2, _ := strconv.Atoi(qStr[0])
	y2, _ := strconv.Atoi(qStr[1])
	q := Point{x2, y2}

	return p, q
}

func pointsOnLine(p Point, q Point) []Point {
	var points []Point
	if p.x == q.x {
		// vertical
		start := p.y
		end := q.y
		if start > q.y {
			start = q.y
			end = p.y
		}
		for y := start; y <= end; y++ {
			points = append(points, Point{p.x, y})
		}
	} else if p.y == q.y {
		// horizontal
		start := p.x
		end := q.x
		if start > q.x {
			start = q.x
			end = p.x
		}
		for x := start; x <= end; x++ {
			points = append(points, Point{x, p.y})
		}
	} else {
		// 45deg diagonal (guaranteed by problem)
		xDelta := 1
		if p.x > q.x {
			xDelta = -1
		}
		yDelta := 1
		if p.y > q.y {
			yDelta = -1
		}

		for x, y := p.x, p.y; x != q.x && y != q.y; x, y = x+xDelta, y+yDelta {
			points = append(points, Point{x, y})
		}
		points = append(points, q)
	}
	return points
}

func addToMap(counts *map[int]map[int]int, pt Point) {
	r := pt.y
	c := pt.x
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
