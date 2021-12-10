package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day8/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	targetNums := 0
	for _, line := range lines {
		_, outputs := parseInputLine(line)
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				targetNums++
			}
		}
	}

	return targetNums
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day8/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	sumOfSums := 0
	for _, line := range lines {
		signals, outputs := parseInputLine(line)
		sumOfSums += solveOutput(signals, outputs)
	}

	return sumOfSums
}

func parseInputLine(line string) ([]string, []string) {
	halves := strings.Split(line, " | ")
	signals := strings.Split(halves[0], " ")
	outputs := strings.Split(halves[1], " ")

	return signals, outputs
}

func solveOutput(signals []string, outputs []string) int {
	signalsByLen := make(map[int][]string)

	for _, signal := range signals {
		l := len(signal)
		chars := strings.Split(signal, "")
		sort.Strings(chars)
		sortedSignal := strings.Join(chars, "")
		_, ok := signalsByLen[l]
		if !ok {
			signalsByLen[l] = make([]string, 0)
		}
		signalsByLen[l] = append(signalsByLen[l], sortedSignal)
	}

	numbers := make(map[int]string)
	numbers[1] = signalsByLen[2][0]
	numbers[4] = signalsByLen[4][0]
	numbers[7] = signalsByLen[3][0]
	numbers[8] = signalsByLen[7][0]

	numbers[2], numbers[3], numbers[5] = determineFiveCharSignals(signalsByLen[5], numbers[4])
	numbers[0], numbers[6], numbers[9] = determineSixCharSignals(signalsByLen[6], numbers[1], numbers[4])

	outputValue := 0

	for _, output := range outputs {
		outputValue *= 10
		switch len(output) {
		case 2:
			outputValue += 1
		case 3:
			outputValue += 7
		case 4:
			outputValue += 4
		case 5:
			if len(diff(output, numbers[2])) == 0 {
				outputValue += 2
			} else if len(diff(output, numbers[3])) == 0 {
				outputValue += 3
			} else {
				outputValue += 5
			}
		case 6:
			if len(diff(output, numbers[0])) == 0 {
				outputValue += 0
			} else if len(diff(output, numbers[6])) == 0 {
				outputValue += 6
			} else {
				outputValue += 9
			}
		case 7:
			outputValue += 8
		}
	}
	return outputValue
}

/*
determineFiveCharSignals
We can determine the 3 with just the length five inputs
if you calculate the number of different characters it will
determine the 3.
2 has 1 difference vs 3 (symmetrical, 2 has 1 3 doesnt, 3 has 1 two doesnt)
5 has 1 difference vs 3 (Again symmetrical)
2 has 2 differences vs 5 (also symmetrical)

if we determine count of diff from the first in slice with other two
if they match, that is the 3. If they are different, the 3 is which ever
one was compared to and got the single difference

to determine 2 and 5, we use the 4 signal and 3. IF we find teh diff of 4 and 3 (as in which parts of 4 arent in 3)
that gives us the top left line. which should only appear in the 5

returns in order two, three, five
*/
func determineFiveCharSignals(fiveCharSignals []string, four string) (string, string, string) {

	zeroVsOne := len(diff(fiveCharSignals[0], fiveCharSignals[1]))
	zeroVsTwo := len(diff(fiveCharSignals[0], fiveCharSignals[2]))

	var two string
	var three string
	var five string

	if zeroVsOne == zeroVsTwo {
		three = fiveCharSignals[0]
		fiveCharSignals[0] = fiveCharSignals[2]
	} else if zeroVsOne < zeroVsTwo {
		three = fiveCharSignals[1]
		fiveCharSignals[1] = fiveCharSignals[2]
	} else {
		three = fiveCharSignals[2]
	}

	fiveCharSignals = fiveCharSignals[:2]

	topLeftRune := diff(four, three)[0]

	if strings.ContainsRune(fiveCharSignals[0], topLeftRune) {
		five = fiveCharSignals[0]
		two = fiveCharSignals[1]
	} else {
		five = fiveCharSignals[1]
		two = fiveCharSignals[0]
	}

	return two, three, five
}

/*
determineSixCharSignals

Very similar idea to determingFiveCharSignals but simplier
6 is missing 1 rune from each of 1 & 4
0 is missing 1 rune from 4 and 0 from 1
9 is missing 0 runes from both 1 & 4
*/
func determineSixCharSignals(signals []string, one string, four string) (string, string, string) {

	var zero string
	var six string
	var nine string

	for _, signal := range signals {
		missingRunesFromOne := len(diff(one, signal))
		missingRunesFromFour := len(diff(four, signal))

		if missingRunesFromOne == 1 && missingRunesFromFour == 1 {
			six = signal
		} else if missingRunesFromFour == 1 {
			zero = signal
		} else {
			nine = signal
		}
	}

	return zero, six, nine
}

func diff(signalA string, signalB string) []rune {
	var notPresent []rune

	for _, r := range signalA {
		if !strings.ContainsRune(signalB, r) {
			notPresent = append(notPresent, r)
		}
	}

	return notPresent
}
