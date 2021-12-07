package main

import (
	"math"
)

type Board struct {
	values []int
	marks  []bool
}

func (b Board) Play(callouts []int) (int, int) {
	b.marks = make([]bool, len(b.values))
	for i, n := range callouts {
		b.handleCallout(n)
		if b.isWinner() {
			return i + 1, b.calcScore(n)
		}
	}
	return math.MaxInt, -1
}

func (b Board) handleCallout(callout int) {
	for i, v := range b.values {
		if v == callout {
			b.marks[i] = true
		}
	}
}

func (b Board) isWinner() bool {
	return b.rowIsWinner(0) ||
		b.rowIsWinner(1) ||
		b.rowIsWinner(2) ||
		b.rowIsWinner(3) ||
		b.rowIsWinner(4) ||
		b.colIsWinner(0) ||
		b.colIsWinner(1) ||
		b.colIsWinner(2) ||
		b.colIsWinner(3) ||
		b.colIsWinner(4)
}

func (b Board) rowIsWinner(row int) bool {
	for i := 0; i < 5; i++ {
		if !b.marks[(row*5)+i] {
			return false
		}
	}
	return true
}

func (b Board) colIsWinner(col int) bool {
	for i := 0; i < 5; i++ {
		if !b.marks[(i*5)+col] {
			return false
		}
	}
	return true
}

func (b Board) calcScore(lastCalled int) int {
	score := 0
	for i, mark := range b.marks {
		if !mark {
			score += b.values[i]
		}
	}
	return score * lastCalled
}
