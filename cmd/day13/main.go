package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

type Fold struct {
	isVertical bool
	foldPt     int
}

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day13/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	dots, folds := parseInput(lines)

	//printDots(dots, len(dots[0]), len(dots))

	dots = performFold(dots, folds[0])

	//printDots(dots, len(dots[0]), len(dots))

	return countDots(dots, len(dots[0]), len(dots))
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day13/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	dots, folds := parseInput(lines)

	for _, fold := range folds {
		dots = performFold(dots, fold)
	}

	printDots(dots, len(dots[0]), len(dots))

	return countDots(dots, len(dots[0]), len(dots))
}

func performFold(dots [][]int, fold Fold) [][]int {
	newHeight := len(dots)
	newWidth := len(dots[0])

	// we are assuming we are folding at midpoints based on input
	if fold.isVertical {
		newHeight = fold.foldPt
	} else {
		newWidth = fold.foldPt
	}

	folded := make([][]int, newHeight)

	for r := range folded {
		folded[r] = make([]int, newWidth)
	}

	for x := 0; x < newWidth; x++ {
		for y := 0; y < newHeight; y++ {
			if fold.isVertical {
				folded[y][x] = int(math.Max(float64(dots[y][x]), float64(dots[len(dots)-y-1][x])))
			} else {
				folded[y][x] = int(math.Max(float64(dots[y][x]), float64(dots[y][len(dots[0])-x-1])))
			}
		}
	}
	return folded
}

func parseInput(lines []string) ([][]int, []Fold) {

	var ys []int
	max_y := 0

	var xs []int
	max_x := 0

	folds := make([]Fold, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			if strings.HasPrefix(parts[0], "fold") {
				instruction := strings.Replace(parts[0], "fold along ", "", 1)
				instParts := strings.Split(instruction, "=")
				val, _ := strconv.Atoi(instParts[1])
				folds = append(folds, Fold{instParts[0] == "y", val})
			}
		} else {
			x, _ := strconv.Atoi(parts[0])
			xs = append(xs, x)
			if x > max_x {
				max_x = x
			}

			y, _ := strconv.Atoi(parts[1])
			ys = append(ys, y)
			if y > max_y {
				max_y = y
			}
		}
	}

	dots := make([][]int, max_y+1)

	for r := range dots {
		dots[r] = make([]int, max_x+1)
	}

	for i := 0; i < len(xs); i++ {
		dots[ys[i]][xs[i]] = 1
	}

	return dots, folds
}

func countDots(dots [][]int, width int, height int) int {
	sum := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			sum += dots[y][x]
		}
	}
	return sum
}

func printDots(dots [][]int, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dots[y][x] == 1 {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}
