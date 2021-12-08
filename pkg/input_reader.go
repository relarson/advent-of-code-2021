package pkg

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// readLines reads a whole file into memory
// and returns a slice of its lines convert to ints
func ReadLinesOfInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

func ReadCommaDelimInts(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	var vals []int

	stringVals := strings.Split(lines[0], ",")
	for _, str := range stringVals {
		val, _ := strconv.Atoi(str)
		vals = append(vals, val)
	}
	return vals, nil
}
